package internal

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"sync"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

const sshAgentPipeWindows = `\\.\pipe\openssh-ssh-agent`

type SSHForwarderConfig struct {
	Username   string
	SSHHost    string
	RemoteHost string
	LocalHost  string
}

type SSHForwarder struct {
	sshConn       *ssh.Client
	localListener net.Listener
	mu            sync.Mutex
	closed        bool
}

// Start sets up and starts the SSH tunnel
func (f *SSHForwarder) Start(config *SSHForwarderConfig) error {
	sshConfig, agentConn, err := getSSHConfig(config.Username)
	if err != nil {
		return err
	}
	defer agentConn.Close()

	// Establish SSH connection
	sshConn, err := ssh.Dial("tcp", config.SSHHost, sshConfig)
	if err != nil {
		return err
	}

	// Start local listener
	localListener, err := net.Listen("tcp", config.LocalHost)
	if err != nil {
		sshConn.Close()
		return err
	}

	// Save references for cleanup
	f.mu.Lock()
	f.sshConn = sshConn
	f.localListener = localListener
	f.closed = false
	f.mu.Unlock()

	// Accept and handle connections
	go f.handleConnections(config.RemoteHost)

	return nil
}

// handleConnections forwards connections from the local listener to the remote host
func (f *SSHForwarder) handleConnections(remoteHost string) {
	for {
		f.mu.Lock()
		if f.closed {
			f.mu.Unlock()
			return
		}
		f.mu.Unlock()

		localConn, err := f.localListener.Accept()
		if err != nil {
			log.Printf("Failed to accept local connection: %v", err)
			continue
		}

		go func() {
			defer localConn.Close()

			remoteConn, err := f.sshConn.Dial("tcp", remoteHost)
			if err != nil {
				log.Printf("Failed to connect to remote host: %v", err)
				return
			}
			defer remoteConn.Close()

			// Forward traffic in both directions
			go io.Copy(remoteConn, localConn)
			io.Copy(localConn, remoteConn)
		}()
	}
}

// Close forcefully shuts down the SSH connection and local listener
func (f *SSHForwarder) Close() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.closed {
		return
	}

	if f.localListener != nil {
		f.localListener.Close()
	}

	if f.sshConn != nil {
		f.sshConn.Close()
	}

	f.closed = true
	log.Println("SSH forwarder has been closed.")
}

// getSSHConfig creates SSH client configuration using the Windows SSH agent
func getSSHConfig(username string) (*ssh.ClientConfig, net.Conn, error) {
	agentClient, conn, err := NewSSHAgent()
	if err != nil {
		return nil, nil, err
	}

	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agentClient.Signers),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Replace with proper host key verification
	}, conn, nil
}

// NewSSHAgent connects to the platform-specific SSH agent
func NewSSHAgent() (agent.Agent, net.Conn, error) {
	switch runtime.GOOS {
	case "windows":
		return windowsSSHAgent()
	default: // Linux/Unix
		authSock := os.Getenv("SSH_AUTH_SOCK")
		if authSock == "" {
			return nil, nil, fmt.Errorf("SSH_AUTH_SOCK is not set")
		}
		conn, err := net.Dial("unix", authSock)
		if err != nil {
			return nil, nil, errors.New("could not connect to Unix SSH agent")
		}
		return agent.NewClient(conn), conn, nil
	}
}

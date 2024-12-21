//go:build unix

package internal

import (
	"errors"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh/agent"
)

func GetSSHAgent() (agent.Agent, net.Conn, error) {
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

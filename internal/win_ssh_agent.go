//go:build windows
package internal

import (
	"errors"
	"net"

	"github.com/Microsoft/go-winio"
	"golang.org/x/crypto/ssh/agent"
)

func windowsSSHAgent() (agent.Agent, net.Conn, error) {
	conn, err := winio.DialPipe(sshAgentPipeWindows, nil)
	if err != nil {
		return nil, nil, errors.New("could not connect to Windows SSH agent")
	}
	return agent.NewClient(conn), conn, nil
}

package internal

import (
	"context"
	"fmt"
	"strings"
)

type Handler struct {
	ctx           context.Context
	activeTunnels map[string]*SSHForwarder
}

func NewHandler() *Handler {
	return &Handler{
		activeTunnels: make(map[string]*SSHForwarder),
	}
}

func (h *Handler) Startup(ctx context.Context) {
	h.ctx = ctx
}

func (h *Handler) GetActiveTunnels() []string {
	// get ids from map and return
	ids := make([]string, 0, len(h.activeTunnels))
	for id := range h.activeTunnels {
		ids = append(ids, id)
	}
	return ids
}

func (h *Handler) Connect(tunnel Tunnel, sshPrivateKeyPath string) error {
	sshForwarder := &SSHForwarder{}
	//check ssh port is given in tunnel.Host if not add default ssh port
	if !strings.Contains(tunnel.Host, ":") {
		tunnel.Host = tunnel.Host + ":22"
	}


	err := sshForwarder.Start(&SSHForwarderConfig{
		Username:   tunnel.Username,
		SSHHost:    tunnel.Host,
		RemoteHost: fmt.Sprintf("%s:%d", tunnel.Target, tunnel.RemotePort),
		LocalHost:  fmt.Sprintf("localhost:%d", tunnel.LocalPort),
	})
	if err != nil {
		return err
	}

	h.addTunnel(tunnel.ID, sshForwarder)
	return nil
}

func (h *Handler) Disconnect(tunnel *Tunnel) {
	h.activeTunnels[tunnel.ID].Close()
	h.removeTunnel(tunnel.ID)
}

func (h *Handler) addTunnel(id string, tunnel *SSHForwarder) {
	h.activeTunnels[id] = tunnel
}

func (h *Handler) removeTunnel(id string) {
	delete(h.activeTunnels, id)
}

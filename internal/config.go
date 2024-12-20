package internal

type Config struct {
	Username          string   `json:"username"`
	Tunnels           []Tunnel `json:"tunnels"`
	SSHPrivateKeyPath string   `json:"sshPrivateKeyPath"`
}

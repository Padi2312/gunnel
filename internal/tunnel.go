package internal

type Tunnel struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Host       string `json:"host"`
	Target     string `json:"target"`
	LocalPort  int    `json:"srcPort"`
	RemotePort int    `json:"destPort"`
}

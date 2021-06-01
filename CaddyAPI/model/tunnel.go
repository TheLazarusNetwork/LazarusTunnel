package model

type Tunnel struct {
	Name      string `json:"name"`
	Port      string `json:"port"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
}

type Tunnels struct {
	Tunnels []Tunnel `json:"tunnels"`
}
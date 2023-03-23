package models

// Device holds the fields the represent a
// network device.
type Device struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	CLLI     string `json:"clli"`
}

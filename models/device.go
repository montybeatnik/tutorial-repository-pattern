package models

// NewDeviceRequest holds the fields to request
// the creation of a new device.
type NewDeviceRequest struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	CLLI     string `json:"clli"`
}

// Device holds the fields the represent a
// network device.
type Device struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	CLLI     string `json:"clli"`
}

package models

type Device struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

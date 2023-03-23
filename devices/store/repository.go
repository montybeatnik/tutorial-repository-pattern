package store

import (
	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
)

// Repository outlines the methods to interact with the device store.
type Repository interface {
	NewDevice(newDevice models.NewDeviceRequest) (models.Device, error)
	GetDeviceByIP(ip string) (models.Device, error)
}

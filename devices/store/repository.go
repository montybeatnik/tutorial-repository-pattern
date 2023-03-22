package store

import (
	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
)

type Repository interface {
	NewDevice(device models.Device) error
	GetDeviceByIP(ip string) (models.Device, error)
}

package store

import (
	"github.com/montybeatnik/tutorials/repository-pattern/models"
)

type Repository interface {
	NewDevice(device models.Device) error
	GetDeviceByIP(ip string) (models.Device, error)
}

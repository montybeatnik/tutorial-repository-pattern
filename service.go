package devices

import (
	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/store"
)

// Service exposes the methods to interact with the store.
type Service interface {
	NewDevice(newDevice models.NewDeviceRequest) (models.Device, error)
	GetDeviceByIP(ip string) (models.Device, error)
}

// service is the concrete implementation of the Service interface.
type service struct {
	repo store.Repository
}

// NewService is a factory function that spins up an instance of our service.
func NewService(repo store.Repository) *service {
	return &service{repo: repo}
}

// NewDevice calls into the repo asking to add a new device.
func (s *service) NewDevice(newDevice models.NewDeviceRequest) (models.Device, error) {
	return s.repo.NewDevice(newDevice)
}

// GetDeviceByIP calls into the repo asking to retrieve a device
// given the specified ip address.
func (s *service) GetDeviceByIP(ip string) (models.Device, error) {
	return s.repo.GetDeviceByIP(ip)
}

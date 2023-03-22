package devices

import (
	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/store"
)

type Service interface {
	StoreDevice(device models.Device) error
	GetDeviceByIP(ip string) (models.Device, error)
}

type service struct {
	repo store.Repository
}

func NewService(repo store.Repository) *service {
	return &service{repo: repo}
}

func (s *service) StoreDevice(device models.Device) error {
	return s.repo.NewDevice(device)
}

func (s *service) GetDeviceByIP(ip string) (models.Device, error) {
	return s.repo.GetDeviceByIP(ip)
}

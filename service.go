package repositorytut

import (
	"github.com/montybeatnik/tutorials/repository-pattern/models"
	"github.com/montybeatnik/tutorials/repository-pattern/store"
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

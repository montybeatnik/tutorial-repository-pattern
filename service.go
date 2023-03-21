package repository

import "github.com/montybeatnik/tutorials/repository-pattern/models"

type Service interface {
	StoreDevice(device models.Device) error
	GetDeviceByIP(ip string) (models.Device, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) StoreDevice(device models.Device) error {
	return s.repo.StoreDevice(device)
}

func (s *service) GetDeviceByIP(ip string) (models.Device, error) {
	return s.repo.GetDeviceByIP(ip)
}

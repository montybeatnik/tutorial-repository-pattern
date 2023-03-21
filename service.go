package repository

type Service interface {
	StoreDevice(device Device) error
	GetDeviceByIP(ip string) (Device, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) StoreDevice(device Device) error {
	return s.repo.StoreDevice(device)
}

func (s *service) GetDeviceByIP(ip string) (Device, error) {
	return s.repo.GetDeviceByIP(ip)
}

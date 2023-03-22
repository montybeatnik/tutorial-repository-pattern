package store

import (
	"errors"

	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
)

type InMemRepo struct {
	store map[int]models.Device
}

var count = 0

func NewInMemRepo() *InMemRepo {
	store := make(map[int]models.Device)
	return &InMemRepo{store: store}
}

func (mr *InMemRepo) NewDevice(device models.Device) error {
	count++
	device.ID = count
	mr.store[count] = device
	return nil
}

func (mr *InMemRepo) GetDeviceByIP(ip string) (models.Device, error) {
	var match int
	for pk, dev := range mr.store {
		if dev.IP == ip {
			match = pk
		}
	}
	if match == 0 {
		return models.Device{}, errors.New("device not found")
	}
	return mr.store[match], nil
}

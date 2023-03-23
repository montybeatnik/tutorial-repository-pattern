package store

import (
	"errors"

	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
)

// InMemRepo stores devices in a map for quick retrieval
// allowing for model changes with minimal effort.
type InMemRepo struct {
	store map[int]models.Device
}

// Starting the PK count at zero using the nil value
// of an int.
var count int

// NewInMemRepo is a factory function that stands up
// our in-memory device store.
func NewInMemRepo() *InMemRepo {
	store := make(map[int]models.Device)
	return &InMemRepo{store: store}
}

// NewDevice adds a device to the map.
func (mr *InMemRepo) NewDevice(newDevice models.NewDeviceRequest) (models.Device, error) {
	count++
	device := mapDeviceAttrs(newDevice)
	device.ID = count
	mr.store[count] = device
	return device, nil
}

// GetDeviceByIP walks the map looking for an occurence
// of the given IP.
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

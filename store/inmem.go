package store

import (
	"log"

	"github.com/montybeatnik/tutorials/repository-pattern/models"
)

type InMemRepo struct {
	store map[int]models.Device
}

var count = 0

func NewInMemRepo() *InMemRepo {
	store := make(map[int]models.Device)
	store[count] = models.Device{Hostname: "hostname1", IP: "1.1.1.1"}
	return &InMemRepo{store: store}
}

func (mr *InMemRepo) StoreDevice(device models.Device) error {
	count++
	mr.store[count] = device
	log.Println(mr.store)
	return nil
}

func (mr *InMemRepo) GetDeviceByIP(ip string) (models.Device, error) {
	var match int
	for pk, dev := range mr.store {
		if dev.IP == ip {
			match = pk
		}
	}
	return mr.store[match], nil
}

package repository

import "log"

type Repository interface {
	StoreDevice(device Device) error
	GetDeviceByIP(ip string) (Device, error)
}

type InMemRepo struct {
	store map[int]Device
}

var count = 0

func NewInMemRepo() *InMemRepo {
	store := make(map[int]Device)
	store[count] = Device{Hostname: "hostname1", IP: "1.1.1.1"}
	return &InMemRepo{store: store}
}

func (mr *InMemRepo) StoreDevice(device Device) error {
	count++
	mr.store[count] = device
	log.Println(mr.store)
	return nil
}

func (mr *InMemRepo) GetDeviceByIP(ip string) (Device, error) {
	var match int
	for pk, dev := range mr.store {
		if dev.IP == ip {
			match = pk
		}
	}
	return mr.store[match], nil
}

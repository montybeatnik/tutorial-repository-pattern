package main

import (
	"log"

	"github.com/montybeatnik/tutorials/repository-pattern/devices"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/store"
)

func main() {
	// Create the in-mem devices store using the repo.
	repo := store.NewInMemRepo()
	// Wire up the repo to the service.
	svc := devices.NewService(repo)
	// Put together a device.
	newDevice := models.NewDeviceRequest{Hostname: "test3", IP: "3.3.3.3", CLLI: "someclli"}
	// Feed that device into the service.
	dev, err := svc.NewDevice(newDevice)
	if err != nil {
		log.Println(err)
	}
	log.Println("created device successfully, the PK is", dev.ID)
	// Retrieve that device from the service.
	dev, err = svc.GetDeviceByIP("3.3.3.3")
	if err != nil {
		log.Println(err)
	}
	// Look at your handy work.
	log.Printf("%+v\n", dev)
}

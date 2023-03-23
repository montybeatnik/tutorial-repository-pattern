package main

import (
	"fmt"
	"log"
	"os"

	"github.com/montybeatnik/tutorials/repository-pattern/devices"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/store"
)

func main() {
	// Create the in-mem devices store using the repo.
	dsn := os.Getenv("DSN")
	repo, err := store.NewPGRepo(dsn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Wire up the repo to the service.
	svc := devices.NewService(repo)
	// Put together a device.
	newDevice := models.Device{
		Hostname: "test3",
		IP:       "3.3.3.3",
		CLLI:     "someclli",
	}
	// Feed that device into the service.
	if err := svc.NewDevice(newDevice); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("blah")
	// Retrieve that device from the service.
	dev, err := svc.GetDeviceByIP(newDevice.IP)
	if err != nil {
		log.Println(err)
	}
	// Look at your handy work.
	log.Printf("%+v\n", dev)
}

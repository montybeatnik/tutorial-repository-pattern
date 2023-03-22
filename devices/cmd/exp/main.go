package main

import (
	"log"
	"os"

	"github.com/montybeatnik/tutorials/repository-pattern/devices"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/store"
)

func main() {
	// Grab our DSN from env.
	DSN := os.Getenv("DSN")
	// Prepare our PG device Store.
	repo, err := store.NewPGRepo(DSN)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Wire up the PG store to our service.
	svc := devices.NewService(repo)
	// Put together a device to "toss it on the shelf" of our PG device store.
	newDevice := models.Device{Hostname: "test3", IP: "3.3.3.3"}
	// Insert the device into the store.
	if err := svc.NewDevice(newDevice); err != nil {
		log.Println(err)
	}
	// Grab it from the "shelf" of our PG store.
	dev, err := svc.GetDeviceByIP("3.3.3.3")
	if err != nil {
		log.Println(err)
	}
	// Relish in the glory of your effort.
	log.Printf("%+v\n", dev)
}

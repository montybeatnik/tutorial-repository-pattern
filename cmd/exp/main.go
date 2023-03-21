package main

import (
	"log"

	"github.com/montybeatnik/tutorials/repository-pattern"
	"github.com/montybeatnik/tutorials/repository-pattern/models"
)

func main() {
	// repo := repository.NewInMemRepo()
	DSN := "postgres://postgres:password@localhost:5432/device_inventory?sslmode=disable"
	repo := repository.NewPGRepo(DSN)
	svc := repository.NewService(repo)
	newDevice := models.Device{Hostname: "test3", IP: "3.3.3.3"}
	if err := svc.StoreDevice(newDevice); err != nil {
		log.Println(err)
	}
	dev, err := svc.GetDeviceByIP("3.3.3.3")
	if err != nil {
		log.Println(err)
	}
	log.Println(dev)
}

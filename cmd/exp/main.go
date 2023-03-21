package main

import (
	"log"

	"github.com/montybeatnik/tutorials/repository-pattern"
	"github.com/montybeatnik/tutorials/repository-pattern/models"
)

func main() {
	repo := repository.NewInMemRepo()
	svc := repository.NewService(repo)
	newDevice := models.Device{Hostname: "test3", IP: "3.3.3.3"}
	svc.StoreDevice(newDevice)
	dev, err := svc.GetDeviceByIP("3.3.3.3")
	if err != nil {
		log.Println(err)
	}
	log.Println(dev)
}

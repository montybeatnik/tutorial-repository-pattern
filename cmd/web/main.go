package main

import (
	"log"
	"net/http"

	repositorytut "github.com/montybeatnik/tutorials/repository-pattern"
	"github.com/montybeatnik/tutorials/repository-pattern/store"
)

func main() {
	port := ":9080"
	repo := store.NewInMemRepo()
	svc := repositorytut.NewService(repo)
	// newDevice := repository.Device{Hostname: "test3", IP: "3.3.3.3"}
	// svc.StoreDevice(newDevice)
	// dev, err := svc.GetDeviceByIP("3.3.3.3")
	svr := repositorytut.NewServer(svc)
	if err := http.ListenAndServe(port, svr.NewMux()); err != nil {
		log.Println(err)
	}
}

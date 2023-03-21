package main

import (
	"log"
	"net/http"

	"github.com/montybeatnik/tutorials/repository"
)

func main() {
	port := ":9080"
	repo := repository.NewInMemRepo()
	svc := repository.NewService(repo)
	// newDevice := repository.Device{Hostname: "test3", IP: "3.3.3.3"}
	// svc.StoreDevice(newDevice)
	// dev, err := svc.GetDeviceByIP("3.3.3.3")
	svr := repository.NewServer(svc)
	if err := http.ListenAndServe(port, svr.NewMux()); err != nil {
		log.Println(err)
	}
}

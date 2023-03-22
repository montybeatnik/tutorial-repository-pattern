package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/montybeatnik/tutorials/repository-pattern/devices"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/store"
)

func main() {

	// =========================================================================
	// Stand up a debug muxer for observability.
	deubgMux := devices.DebugStandardLibraryMux()

	// Fire up a web server in the background, exposing the debug methods.
	go func() {
		if err := http.ListenAndServe(":4001", deubgMux); err != nil {
			log.Println(err)
		}
	}()

	port := flag.String("port", "9080", "port on which to listen for incoming requests")
	flag.Parse()
	repo := store.NewInMemRepo()
	svc := devices.NewService(repo)
	svr := devices.NewServer(svc)
	log.Printf("firing up server on %v", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *port), svr.NewMux()); err != nil {
		log.Println(err)
	}
}

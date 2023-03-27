package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/montybeatnik/tutorials/repository-pattern/devices"
	"github.com/montybeatnik/tutorials/repository-pattern/devices/store"
	"github.com/pkg/profile"
)

func main() {

	// =========================================================================
	// Collect and store a memory profile. Although you can modify the self-referncial
	// funcs in the start func to change the profile type (or add a trace).
	//
	// Uncomment for performance gain.
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	// =========================================================================
	// Stand up a debug muxer for observability.
	// deubgMux := DebugStandardLibraryMux()

	// Fire up a web server in the background, exposing the debug methods.
	// go func() {
	// 	if err := http.ListenAndServe(":4001", deubgMux); err != nil {
	// 		log.Println(err)
	// 	}
	// }()

	// grab the port from the CLI or use 8000 by defaul
	port := flag.String("port", "8000", "port on which to listen for incoming requests")
	flag.Parse()
	// stand up the repo
	repo := store.NewInMemRepo()
	// wire the repo into the service
	svc := devices.NewService(repo)
	// build a server
	svr := NewServer(svc)
	// fire up the server
	log.Printf("firing up server on %v", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *port), svr.NewMux()); err != nil {
		log.Println(err)
	}
}

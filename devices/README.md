# Repository Tutorial 
The Repository Tutorial aims to cover the usefuless of the repository pattern. This service exposes an API to interact with Devices. At the moment, a device has a small number of attributes, but this could grow to exponential porportions. 

## Usage 
### Experimenting
There is a cmd/exp/main.go file for testing out ideas and interacting with the service without the need for a server. 

```bash
# From the project root /devices. 
go run cmd/exp/main.go
```

#### Sample code for the exp dir.
##### InMem example 
```go
func main() {
	// Create the in-mem devices store using the repo.
	repo := store.NewInMemRepo()
	// Wire up the repo to the service.
	svc := devices.NewService(repo)
	// Put together a device.
	newDevice := models.Device{Hostname: "test3", IP: "3.3.3.3"}
	// Feed that device into the service.
	if err := svc.NewDevice(newDevice); err != nil {
		log.Println(err)
	}
	// Retrieve that device from the service.
	dev, err := svc.GetDeviceByIP("3.3.3.3")
	if err != nil {
		log.Println(err)
	}
	// Look at your handy work.
	log.Printf("%+v\n", dev)
}
```

##### PG example 
```go
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
```

### Web 
If you're into firing up a full on web server and hitting the API with JSON or to get JSON, use cmd/web/main.go. 

#### Sample code for the web dir. 
```go
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
	// grab the port from the CLI or use 9080 by defaul
	port := flag.String("port", "9080", "port on which to listen for incoming requests")
	flag.Parse()
	// stand up the repo
	repo := store.NewInMemRepo()
	// wire the repo into the service
	svc := devices.NewService(repo)
	// build a server
	svr := devices.NewServer(svc)
	// fire up the server
	log.Printf("firing up server on %v", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *port), svr.NewMux()); err != nil {
		log.Println(err)
	}
}
```

## Curl examples 
```bash
# create a device
curl -X POST localhost:9080/new-device \
    -H 'Content-Type: application/json' \
    -d '{"hostname": "curl_test", "ip": "7.7.7.42", "clli": "someclli"}'
# get a device by ip
curl -X GET localhost:9080/device/7.7.7.42
# if you're into metrics
curl -X GET localhost:4000/debug/vars
```

## Dependencies
You can always just check the go.mod file, but for this project, we're using 1 external dependency: lib/pq. 

```bash
# Which you can just go get: (I love go!)
go get github.com/lib/pq
```

## Env vars
`DSN` 
```bash
export DSN="postgres://postgres:password@localhost:5432/device_inventory?sslmode=disable"
```

## Testing
Sadly, there are no tests. 
package devices

import (
	"encoding/json"
	"expvar"
	"log"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
)

type server struct {
	deviceService Service
	router        *http.ServeMux
}

func (s *server) NewMux() http.Handler {
	s.router.HandleFunc("/new-device", s.handleNewDevice)
	s.router.HandleFunc("/device/", s.handleDeviceByIP)
	s.router.HandleFunc("/ping", s.handlePing)
	return s.router
}

func (s *server) handleNewDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	var device models.Device
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		log.Println("decoding failed", err)
	}
	log.Println("creating a device", device)
	if err := s.deviceService.StoreDevice(device); err != nil {
		log.Println(err)
	}
}

func (s *server) handleDeviceByIP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	var ip string
	if len(strings.Split(r.URL.Path, "/")) > 1 {
		ip = strings.Split(r.URL.Path, "/")[2]
	}
	dev, err := s.deviceService.GetDeviceByIP(ip)
	if err != nil {
		log.Println(err)
		resp := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&resp)
		return
	}
	json.NewEncoder(w).Encode(&dev)
}

func (s *server) handlePing(w http.ResponseWriter, r *http.Request) {
	pong := map[string]string{"msg": "pong"}
	json.NewEncoder(w).Encode(pong)
	return
}

func NewServer(svc Service) *server {
	mux := http.NewServeMux()
	return &server{
		deviceService: svc,
		router:        mux,
	}
}

// StandardLibraryMux registers all the debug routes from the standard library
// into a new mux bypassing the use of the DefaultServerMux. Using the
// DefaultServerMux would be a security risk since a dependency could inject a
// handler into our service without us knowing it.
func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}

package repository

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type server struct {
	deviceService Service
	router        *http.ServeMux
}

func (s *server) NewMux() http.Handler {
	s.router.HandleFunc("/new-device", s.newDevice)
	s.router.HandleFunc("/device/", s.device)
	s.router.HandleFunc("/ping", s.ping)
	return s.router
}

func (s *server) newDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	var device Device
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		log.Println("decoding failed", err)
	}
	log.Println("creating a device", device)
	if err := s.deviceService.StoreDevice(device); err != nil {
		log.Println(err)
	}
}

func (s *server) device(w http.ResponseWriter, r *http.Request) {
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
	}
	json.NewEncoder(w).Encode(&dev)
}

func (s *server) ping(w http.ResponseWriter, r *http.Request) {
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

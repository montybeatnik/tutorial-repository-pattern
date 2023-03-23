package devices

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"
)

type ctxKey int

const key ctxKey = 1

// Values represent state for each request.
type Values struct {
	Now        time.Time
	StatusCode int
}

// SetStatusCode sets the status code back into the context.
func SetStatusCode(ctx context.Context, statusCode int) {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return
	}

	v.StatusCode = statusCode
}

// respond converts a Go value to JSON and sends it to the client.
func respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {

	SetStatusCode(ctx, statusCode)

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// handleNewDevice
func (s *server) handleNewDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	var newDeviceRequest models.NewDeviceRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&newDeviceRequest); err != nil {
		log.Println("decoding failed", err)
	}
	log.Println("creating a device", newDeviceRequest)
	device, err := s.deviceService.NewDevice(newDeviceRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println("device created with pk:", device.ID)
	if err := respond(r.Context(), w, device, http.StatusCreated); err != nil {
		log.Println(err)
		return
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

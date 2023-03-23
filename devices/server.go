package devices

import (
	"expvar"
	"net/http"
	"net/http/pprof"
)

// server defines the dependency the server needs.
type server struct {
	deviceService Service
	router        *http.ServeMux
}

// NewMux defines the routes exposed via the endpoint, mapping
// the routes to handlers.
func (s *server) NewMux() http.Handler {
	s.router.HandleFunc("/new-device", s.handleNewDevice)
	s.router.HandleFunc("/device/", s.handleDeviceByIP)
	s.router.HandleFunc("/ping", s.handlePing)
	return s.router
}

// NewServer is a factory function that injects the router
// and the service to the server.
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

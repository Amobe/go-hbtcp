package admin

import (
	"net/http"
)

type Stats interface {
	ToJson() []byte
}

type AdminServer struct {
	address string
	stats   Stats
}

func NewAdminServer(address string, stats Stats) *AdminServer {
	return &AdminServer{
		address,
		stats,
	}
}

func (a *AdminServer) Start() {
	http.HandleFunc("/stat", a.handleStatRequest)
	http.ListenAndServe(a.address, nil)
}

func (a *AdminServer) handleStatRequest(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	w.Write(a.stats.ToJson())
}

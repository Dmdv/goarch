package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	Router *mux.Router
}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes ...
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Service is running")
	})
}

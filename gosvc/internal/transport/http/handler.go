package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goarch/gosvc/internal/comment"
	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response ...
type Response struct {
	Message string
}

// NewHandler ...
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes ...
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "Service running OK"}); err != nil {
			panic(err)
		}
	})
}

// GetComment ...
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprint(w, "Unable to parse UINT from ID")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprint(w, "Error retreiving comment by ID")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// GetAllComments ...
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprint(w, "Failed to retreive all comments")
	}

	fmt.Fprintf(w, "%+v", comments)
}

// PostComment ...
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprint(w, "Failed to post new comment")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// UpdateComment ...
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})

	if err != nil {
		fmt.Fprint(w, "Error retreiving comment by ID")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// DeleteComment ...
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprint(w, "Unable to parse UINT from ID")
	}

	err = h.Service.DeleteComment(uint(i))

	if err != nil {
		fmt.Fprint(w, "Failed to delete comment by ID")
	}

	fmt.Fprintf(w, "Successfully deleted comment")
}

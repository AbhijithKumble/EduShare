package serverhealth

import (
	"net/http"

	"github.com/AbhijithKumble/EduShare/backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/health", h.healthGet).Methods("GET")
	r.HandleFunc("/health", h.healthPost).Methods("POST")
	r.HandleFunc("/health", h.healthDelete).Methods("DELETE")
	r.HandleFunc("/health", h.healthPatch).Methods("PATCH")
}

func (h *Handler) healthPost(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "Hello this route is ready handle POST request")
}

func (h *Handler) healthPatch(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "Hello this route is ready handle PATCH request")
}

func (h *Handler) healthGet(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "Hello this route is ready handle GET request")
}

func (h *Handler) healthDelete(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "Hello this route is ready handle DELETE request")
}


package verify

import (
	"net/http"
	"net/smtp"

	"github.com/gorilla/mux"
)

var (
	from = "adk1543@yahoo.com"
	msg  = []byte("dummy message")
    hostname = ""
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/verify", h.HandleVerify).Methods("POST")
	r.HandleFunc("/verify", h.HandleCheckVerify).Methods("GET")
}

func (h *Handler) HandleVerify(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleCheckVerify(w http.ResponseWriter, r *http.Request) {
}

package signup

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AbhijithKumble/EduShare/backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/signup", h.HandleSignup).Methods("POST")
}

func (h *Handler) HandleSignup(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(params)

	if err != nil {
		utils.RespondWithJSON(w, 400, "Invalid JSON format")
		return
	}

	//create user
	err = h.store.CreateUser(r.Context(), params.Email, params.Password)

	if err != nil {
		switch err.Error() {
		case "user already exists in database":
			utils.RespondWithError(w, 409, "User already exists")

		default:
			log.Printf("error creating user %v", err)
			utils.RespondWithError(w, 500, "Something went wrong")
		}
		return
	}

	utils.RespondWithJSON(w, 201, "User created Successfully")
}

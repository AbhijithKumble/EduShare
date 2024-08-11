package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AbhijithKumble/EduShare/backend/configs"
	"github.com/AbhijithKumble/EduShare/backend/services/auth"
	"github.com/AbhijithKumble/EduShare/backend/types"
	"github.com/AbhijithKumble/EduShare/backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore 
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/login", h.HandleSignup).Methods("POST")
	r.HandleFunc("/signup", h.HandleSignup).Methods("POST")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    var user types.LoginUserPayload 

    if err := utils.ParseJSON(r, &user); err != nil  {
        utils.WriteError(w, http.StatusBadRequest,err)
    }
    
    if err := utils.Validate.Struct(user); err != nil {
        errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
    }

    u, err := h.store.GetUserByEmail(r.Context(), user.Email)

    if err!= nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
    }

    if !auth.ComparePassword(u.Password, user.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

    secret := []byte(configs.Envs.JWT_SECRET) 

    token, err := auth.CreateJWT(secret, u.UserID)

    if err!= nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return 
    }

    utils.WriteJSON(w, http.StatusOK , map[string]string {"token": token})
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
		utils.WriteError(w, 400, fmt.Errorf("Invalid JSON format"))
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

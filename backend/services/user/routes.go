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
	r.HandleFunc("/login", h.HandleLogin).Methods("POST")
	r.HandleFunc("/signup", h.HandleSignup).Methods("POST")
	r.HandleFunc("/verifyemail/${userID}", h.HandleVerifyEmail).Methods("POST") //middleware
	r.HandleFunc("/forgotpassword/${userID}", h.HandleForgotPassword).Methods("POST")
	r.HandleFunc("/resetpassword/${userID}", h.HandleResetPassword).Methods("POST") //create verifyemailmiddleware

	//user routes
	//change it to home
	r.HandleFunc("/users/{userID}", auth.WithJWT(h.HandleGetUser, h.store)).Methods("GET")
	r.HandleFunc("/users/{userID}/addDetails", auth.WithJWT(h.HandleGetUser, h.store)).Methods("GET")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err, _ := h.store.GetUserByEmail(r.Context(), user.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePassword(u.Password, user.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	isVerified := u.IsVerified

	if !isVerified {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("User account not verified"))
		return
	}

	secret := []byte(configs.Envs.JWT_SECRET)

	token, err := auth.CreateJWT(secret, u.UserID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
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
	var user types.UserAcc
	user.Email = params.Email
	user.Password = params.Password
	err = h.store.CreateUser(r.Context(), user)

	if err != nil {
		switch err.Error() {
		case "user is not verified":
			utils.WriteError(w, 400, fmt.Errorf("User is not verified"))

		case "user already exists in database":
			utils.WriteError(w, 409, fmt.Errorf("User already exists"))
    
    case "error sending mail":
			utils.WriteError(w, 500, fmt.Errorf("Something went wrong"))

		default:
			log.Printf("error creating user -> %v", err)
			utils.WriteError(w, 500, fmt.Errorf("Something went wrong"))
		}
		return
	}
  
	utils.WriteJSON(w, 201, "User created Successfully")
}

func (h *Handler) HandleForgotPassword(w http.ResponseWriter, r *http.Request) {
	var payload types.ForgotPasswordPayLoad

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&payload)
	email := payload.Email

	_, err, code := h.store.GetUserByEmail(r.Context(), email)

	if err != nil {
		utils.WriteError(w, code, err)
	}

	utils.WriteJSON(w, http.StatusOK, "User found")
}

func (h *Handler) HandleResetPassword(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleVerifyEmail(w http.ResponseWriter, r *http.Request) {
//  var userID uuid.UUID
  
 // decoder := json.NewDecoder(r.Body)

}

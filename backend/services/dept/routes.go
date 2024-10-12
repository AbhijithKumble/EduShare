package dept

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbhijithKumble/EduShare/backend/types"
	"github.com/AbhijithKumble/EduShare/backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.DeptStore
}

func NewHandler(store types.DeptStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/getdepts", h.HandleGetDepts).Methods("GET")

	// admin route
	r.HandleFunc("/adddepts", h.HandleAddDepts).Methods("POST")
}

func (h *Handler) HandleGetDepts(w http.ResponseWriter, r *http.Request) {
	courses, err := h.store.GetDepts(r.Context())

	if err != nil {
		utils.WriteError(w, 500, fmt.Errorf("Failed to retreive courses"))
		return
	}

	utils.WriteJSON(w, 200, courses)
}

func (h *Handler) HandleAddDepts(w http.ResponseWriter, r *http.Request) {
	var dept types.DeptPayload

	Decoder := json.NewDecoder(r.Body)
	Decoder.DisallowUnknownFields()
	err := Decoder.Decode(&dept)
	defer r.Body.Close()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if dept.DeptCode == "" || dept.DeptName == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Does not contain all required fields"))
		return
	}

	err = h.store.CreateDepts(r.Context(), dept)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "Dept created successfully")
	return
}

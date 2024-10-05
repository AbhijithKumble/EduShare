package courses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbhijithKumble/EduShare/backend/types"
	"github.com/AbhijithKumble/EduShare/backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CourseStore
}

func NewHandler(store types.CourseStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/getcourses", h.HandleGetCourses).Methods("GET")

	// user route
	r.HandleFunc("/favourites", h.HandleGetFavouriteCourse).Methods("GET")

	// admin route
	r.HandleFunc("/addcourses", h.HandleGetFavouriteCourse).Methods("POST")
}

func (h *Handler) HandleGetCourses(w http.ResponseWriter, r *http.Request) {
  courses, err := h.store.GetCourses(r.Context())

  if err!=nil {
    utils.WriteError(w, 500, fmt.Errorf("Failed to retreive courses"))
    return 
  }
  
  utils.WriteJSON(w, 200, courses)
}

func (h *Handler) HandleGetFavouriteCourse(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) HandleAddCourses(w http.ResponseWriter, r *http.Request) {
  var course types.Course

  err := json.NewDecoder(r.Body).Decode(&course)
  defer r.Body.Close()

  if err!= nil {
    utils.WriteError(w, http.StatusBadRequest, err)
    return
  }

}


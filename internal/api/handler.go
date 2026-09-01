package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/davenfinch/followmyjourney-profile-service/internal/model"
	"github.com/davenfinch/followmyjourney-profile-service/internal/service"
)

type Handler struct {
	service *service.ProfileService
}

func NewHandler(s *service.ProfileService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/healthz", h.health).Methods("GET")
	r.HandleFunc("/profiles", h.createProfile).Methods("POST")
	r.HandleFunc("/profiles/{id}", h.getProfile).Methods("GET")
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *Handler) createProfile(w http.ResponseWriter, r *http.Request) {
	var p model.Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateProfile(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *Handler) getProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	p, err := h.service.GetProfile(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

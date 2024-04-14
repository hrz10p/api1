package handlers

import (
	"ass1/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	s *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create-module", h.CreateModuleInfoHandler)
	mux.HandleFunc("/get-module", h.GetModuleInfoHandler)
	mux.HandleFunc("/edit-module", h.EditModuleInfoHandler)
	mux.HandleFunc("/delete-module", h.DeleteModuleInfoHandler)

	return mux
}

type Response struct {
	Message string `json:"message"`
}

func (h *Handler) errorResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	errResponse := Response{
		Message: http.StatusText(statusCode),
	}
	json.NewEncoder(w).Encode(errResponse)
}

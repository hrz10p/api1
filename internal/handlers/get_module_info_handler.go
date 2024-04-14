package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) GetModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorResponse(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest)
		return
	}
	moduleInfo, err := h.s.GetModuleInfo(id)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(moduleInfo)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

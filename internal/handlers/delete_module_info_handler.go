package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.errorResponse(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest)
		return
	}
	if err := h.s.DeleteModuleInfo(id); err != nil {
		h.errorResponse(w, http.StatusInternalServerError)
		return
	}

	response := Response{
		Message: "Module info deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

package handlers

import (
	"ass1/internal/data"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) EditModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		h.errorResponse(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest)
		return
	}
	var updatedInfo data.ModuleInfo
	if err := json.NewDecoder(r.Body).Decode(&updatedInfo); err != nil {
		h.errorResponse(w, http.StatusBadRequest)
		return
	}
	if id != updatedInfo.ID {
		h.errorResponse(w, http.StatusBadRequest)
		return
	}
	moduleInfo, err := h.s.UpdateModuleInfo(updatedInfo)
	if err != nil {
		log.Println(err.Error())
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

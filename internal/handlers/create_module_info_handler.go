package handlers

import (
	"ass1/internal/data"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) CreateModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var createModuleInfo data.ModuleInfo
	if err := json.NewDecoder(r.Body).Decode(&createModuleInfo); err != nil {
		h.errorResponse(w, http.StatusBadRequest)
		return
	}

	moduleInfo, err := h.s.CreateModuleInfo(createModuleInfo)
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

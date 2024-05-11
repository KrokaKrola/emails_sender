package handlers

import "net/http"

type HealthHandlers struct {
}

func NewHealthHandlers() *HealthHandlers {
	return &HealthHandlers{}
}

func (h *HealthHandlers) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

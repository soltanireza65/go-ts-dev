package handlers

import "net/http"

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Excute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

package otp

import (
	"eschool/rest/middlewares"
	"net/http"
)

// RegisterRoutes sets up OTP routes
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /users/send-otp",
		manager.With(http.HandlerFunc(h.SendOTP)),
	)
}

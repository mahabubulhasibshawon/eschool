package user

import (
	// "crypto/tls"
	"eschool/rest/handlers/otp"
	"eschool/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	otpHandler  *otp.Handler
}

func NewHandler(middlewares *middlewares.Middlewares, otpHandler *otp.Handler) *Handler {
	return &Handler{
		middlewares: middlewares,
		otpHandler:  otpHandler,
	}
}

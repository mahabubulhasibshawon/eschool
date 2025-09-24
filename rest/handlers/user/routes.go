package user

import (
	"eschool/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		),
	)
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)
	mux.Handle(
		"POST /users/refresh",
		manager.With(
			http.HandlerFunc(h.Refresh),
		),
	)
	
}
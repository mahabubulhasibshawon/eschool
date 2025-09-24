package course

import (
	"eschool/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /courses",
		manager.With(
			http.HandlerFunc(h.GetCourses),
		),
	)
	mux.Handle(
		"POST /courses",
		manager.With(
			http.HandlerFunc(h.CreateCourse),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"GET /courses/{id}",
		manager.With(
			http.HandlerFunc(h.GetCourse),
		),
	)
	mux.Handle(
		"PUT /courses/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateCourse),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"DELETE /courses/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteCourse),
			h.middlewares.AuthenticateJWT,
		),
	)
}

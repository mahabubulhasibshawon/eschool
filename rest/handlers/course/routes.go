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
		),
	)
// 	mux.Handle(
// 		"GET /courses/{id}",
// 		manager.With(
// 			http.HandlerFunc(h.GetCourse),
// 		),
// 	)
// 	mux.Handle(
// 		"DELETE /courses/{id}",
// 		manager.With(
// 			http.HandlerFunc(h.DeleteCourse),
// 		),
// 	)
}

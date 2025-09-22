package course

import (
	"eschool/database"
	"eschool/util"
	"log"
	"net/http"
)

func (h *Handler) GetCourses(w http.ResponseWriter, r *http.Request) {
	var courses []database.Course

	query := `SELECT * FROM courses`

	err := h.middlewares.DB.Select(&courses, query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}
	util.SendData(w, courses, http.StatusOK)

}
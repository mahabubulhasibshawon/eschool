package course

import (
	"eschool/database"
	"eschool/util"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {
	courseID := r.PathValue("id")

	cId, err := strconv.Atoi(courseID)
	if err != nil {
		http.Error(w, "Please give me valid course id", 400)
		return
	}

	// course := database.GetCourse(cId)
	var courses database.Course
	query := `select * from courses where id=$1`
	err = h.middlewares.DB.Get(&courses, query, cId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch course", http.StatusInternalServerError)
		return
	}

	util.SendData(w, courses, 200)
}

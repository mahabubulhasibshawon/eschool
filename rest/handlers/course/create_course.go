package course

import (
	"encoding/json"
	"eschool/database"
	"eschool/util"
	"log"
	"net/http"
)

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var newCourse database.Course
	err := json.NewDecoder(r.Body).Decode(&newCourse)
	if err != nil {
		log.Println(err)
		http.Error(w, "Please send a valid JSON body", http.StatusBadRequest)
		return
	}
	var courses database.Course
	query := `
        INSERT INTO courses
        (title, instructor, description, category)
        VALUES
        ($1, $2, $3, $4)
        RETURNING id, title, instructor, description, category
    `

	err = h.middlewares.DB.Get(&courses, query, newCourse.Title,newCourse.Instructor, newCourse.Description, newCourse.Category)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to store course", http.StatusInternalServerError)
		return
	}

	// createdCourse := database.StoreCourse(newCourse)
	util.SendData(w, courses, http.StatusCreated)

}

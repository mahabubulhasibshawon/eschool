package course

import (
	"encoding/json"
	"eschool/database"
	"eschool/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	courseID := r.PathValue("id")

	cId, err := strconv.Atoi(courseID)
	if err != nil {
		http.Error(w, "Please give me valid course id", 400)
		return
	}

	var course database.Course
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&course)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please send a valid JSON body", 400)
		return
	}

	// SQL query as a variable
	query := `
		UPDATE courses
		SET title = $1,
			instructor_id = $2,
			description = $3,
			category = $4
		WHERE id = $5
		RETURNING id, title, instructor_id, description, category
	`

	var updatedCourse database.Course

	err = h.middlewares.DB.Get(&updatedCourse, query, course.Title, course.InstructorID, course.Description, course.Category, cId)

	if err != nil {
		http.Error(w, "Failed to update course: "+err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Successfully updated course", 201)
}

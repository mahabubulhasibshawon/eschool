package course

import (
	"encoding/json"
	"eschool/database"
	"eschool/util"
	"log"
	"net/http"
)

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	// Get user_id from context (set by AuthenticateJWT)
	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		http.Error(w, "Unauthorized: No user ID", http.StatusUnauthorized)
		return
	}

	// Parse input
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Data!", http.StatusBadRequest)
		return
	}

	// Validate input
	if input.Title == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Create course
	course := database.Course{
		Title:        input.Title,
		InstructorID: userID,
		Description:  input.Description,
		Category:     input.Category,
	}
	query := `
        INSERT INTO courses (title, instructor_id, description, category, created_at)
        VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
        RETURNING id, title, instructor_id, description, category, created_at
    `
	err = h.middlewares.DB.Get(&course, query, course.Title, course.InstructorID, course.Description, course.Category)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to store course", http.StatusInternalServerError)
		return
	}

	// Send response
	util.SendData(w, course, http.StatusCreated)
}
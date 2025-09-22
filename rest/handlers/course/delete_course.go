package course

import (
	"eschool/util"

	"net/http"
	"strconv"
)

func (h *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	courseId := r.PathValue("id")

	cId, err := strconv.Atoi(courseId)
	if err != nil {
		http.Error(w, "Please give me valid course id", 400)
		return
	}

	query := `DELETE FROM courses WHERE id = $1`
	res, err := h.middlewares.DB.Exec(query, cId)
	if err != nil {
		http.Error(w, "Failed to delete course"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any row was affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Could not verify deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No course found with this ID", http.StatusNotFound)
		return
	}

	util.SendData(w, map[string]any{
		"message": "Course deleted successfully",
		"id":      cId,
	}, http.StatusOK)
}

package course

// import (
// 	"eschool/database"
// 	"eschool/util"
// 	"net/http"
// 	"strconv"
// )

// func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {
// 	courseID := r.PathValue("id")

// 	cId, err := strconv.Atoi(courseID)
// 	if err != nil {
// 		http.Error(w, "Please give me valid course id", 400)
// 		return
// 	}

// 	course := database.GetCourse(cId)

// 	query := `select * from courses where id=$1`
// 	err := h.middlewares.DB.Get(query, &newCourse, id)

// 	if course == nil {
// 		util.SendError(w, 404, "course not found!")
// 		return
// 	}

// 	util.SendData(w, course, 200)
// }

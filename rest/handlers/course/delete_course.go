package course

// import (
// 	"eschool/util"
// 	"eschool/database"

// 	"net/http"
// 	"strconv"
// )

// func (h *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
// 	courseId := r.PathValue("id")

// 	cId, err := strconv.Atoi(courseId)
// 	if err != nil {
// 		http.Error(w, "Please give me valid product id", 400)
// 		return
// 	}

// 	database.DeleteCourse(cId)

// 	util.SendData(w, "Successfully deleted product!", 201)
// }

package course

// import (
// 	"encoding/json"
// 	"eschool/database"
// 	"eschool/util"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )

// func (h *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
// 	courseID := r.PathValue("id")

// 	cId, err := strconv.Atoi(courseID)
// 	if err != nil {
// 		http.Error(w, "Please give me valid course id", 400)
// 		return
// 	}

// 	var newCourse database.Course
// 	decoder := json.NewDecoder(r.Body)
// 	err = decoder.Decode(&newCourse)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Please send a valid JSON body", 400)
// 		return
// 	}

// 	newCourse.ID = cId

// 	database.UpdateCourse(newCourse)

// 	util.SendData(w, "Successfully updated course", 201)
// }

package user

import (
	"encoding/json"
	"eschool/database"
	"eschool/util"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//	func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
//		var newUser database.User
//		decoder := json.NewDecoder(r.Body)
//		err := decoder.Decode(&newUser)
//		if err != nil {
//			fmt.Println(err)
//			http.Error(w, "Invalid Request Data!", http.StatusBadRequest)
//			return
//		}
//		var users database.User
//		query := `
//	        INSERT INTO users
//	        (username, email, password_hash, created_at)
//	        VALUES
//	        ($1, $2, $3, $4)
//	        RETURNING id, username, email, created_at
//	    `
//		err = h.middlewares.DB.Get(&users, query, newUser.Username, newUser.Email, newUser.PasswordHash, newUser.CreatedAt)
//		if err != nil {
//			log.Println(err)
//			http.Error(w, "Failed to store user", http.StatusInternalServerError)
//			return
//		}
//		util.SendData(w, users, 201)
//	}
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Data!",http.StatusBadRequest)
		return
	}
	// Validate input
	if input.Username == "" || input.Email == "" || input.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// create user
	newUser := database.User {
		Username: input.Username,
		Email: input.Email,
		PasswordHash: string(hash),
	}
	var createdUser database.User
	query := `
        INSERT INTO users
        (username, email, password_hash, created_at)
        VALUES
        ($1, $2, $3, CURRENT_TIMESTAMP)
        RETURNING id, username, email, created_at
    `
	err = h.middlewares.DB.Get(&createdUser, query, newUser.Username, newUser.Email, newUser.PasswordHash)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to store user", http.StatusInternalServerError)
		return
	}

	// send response
	util.SendData(w, createdUser, http.StatusCreated)
}

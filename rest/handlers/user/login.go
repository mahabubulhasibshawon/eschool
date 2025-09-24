package user

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"eschool/config"
	"eschool/database"
	"eschool/util"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims matches AuthenticateJWT middleware
type Claims struct {
	UserID    int   `json:"user_id"`
	ExpiresAt int64 `json:"exp"`
}

// GetUserByUsername fetches user by username
func (h *Handler) GetUserByUsername(username string) (database.User, error) {
	var user database.User
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE username = $1`
	err := h.middlewares.DB.Get(&user, query, username)
	return user, err
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requestLogin RequestLogin
	err := json.NewDecoder(r.Body).Decode(&requestLogin)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Data!", http.StatusBadRequest)
		return
	}

	// Validate input
	if requestLogin.Username == "" || requestLogin.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Fetch user
	user, err := h.GetUserByUsername(requestLogin.Username)
	if err != nil || user.ID == 0 {
		log.Println(err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestLogin.Password)); err != nil {
		log.Println(err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	cnf := config.GetConfig()
	accessToken, err := generateAccessToken(user.ID, cnf.JwtSecretKey)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Generate and store refresh token
	refreshToken, err := h.generateRefreshToken(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send response
	util.SendData(w, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, http.StatusOK)
}
func generateAccessToken(userID int, secret string) (string, error) {
	claims := Claims{
		UserID:    userID,
		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
	}
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	headerEnc := base64urlEncode(headerBytes)
	payloadEnc := base64urlEncode(payloadBytes)
	message := headerEnc + "." + payloadEnc
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	signature := base64urlEncode(h.Sum(nil))
	return headerEnc + "." + payloadEnc + "." + signature, nil
}

func base64urlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

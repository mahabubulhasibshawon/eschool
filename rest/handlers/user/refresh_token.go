package user

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"eschool/config"
	"eschool/database"
	"eschool/util"
	"log"
	"net/http"
	"time"
)

type RequestRefresh struct {
	RefreshToken string `json:"refresh_token"`
}

// GetRefreshToken fetches and validates a refresh token
func (h *Handler) GetRefreshToken(token string) (database.RefreshToken, error) {
	var rt database.RefreshToken
	query := `SELECT id, user_id, token, expires_at, revoked, created_at FROM refresh_tokens WHERE token = $1 AND revoked = FALSE`
	err := h.middlewares.DB.Get(&rt, query, token)
	return rt, err
}

// GetUserById fetches user by ID (from prior, defined in handler file)
func (h *Handler) GetUserById(id int) (database.User, error) {
	var user database.User
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE id = $1`
	err := h.middlewares.DB.Get(&user, query, id)
	return user, err
}

// RevokeRefreshToken marks a refresh token as revoked
func (h *Handler) RevokeRefreshToken(token string) error {
	query := `UPDATE refresh_tokens SET revoked = TRUE WHERE token = $1`
	_, err := h.middlewares.DB.Exec(query, token)
	return err
}

// generateRefreshToken creates and stores a refresh token
func (h *Handler) generateRefreshToken(userID int) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := base64.StdEncoding.EncodeToString(b)
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	rt := database.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	query := `
        INSERT INTO refresh_tokens (user_id, token, expires_at)
        VALUES ($1, $2, $3)
    `
	_, err := h.middlewares.DB.Exec(query, rt.UserID, rt.Token, rt.ExpiresAt)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var requestRefresh RequestRefresh
	err := json.NewDecoder(r.Body).Decode(&requestRefresh)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Data!", http.StatusBadRequest)
		return
	}

	// Validate input
	if requestRefresh.RefreshToken == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Fetch and validate refresh token
	rt, err := h.GetRefreshToken(requestRefresh.RefreshToken)
	if err != nil || rt.UserID == 0 || rt.Revoked || rt.ExpiresAt.Before(time.Now()) {
		log.Println(err)
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	// Fetch user
	user, err := h.GetUserById(rt.UserID)
	if err != nil || user.ID == 0 {
		log.Println(err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Revoke old refresh token
	if err := h.RevokeRefreshToken(requestRefresh.RefreshToken); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Generate new access token
	cnf := config.GetConfig()
	newAccessToken, err := generateAccessToken(user.ID, cnf.JwtSecretKey)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Generate and store new refresh token
	newRefreshToken, err := h.generateRefreshToken(user.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send response
	util.SendData(w, map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	}, http.StatusOK)
}

// // generateAccessToken (unchanged from prior)
// func generateAccessToken(userID int, secret string) (string, error) {
// 	claims := Claims{
// 		UserID:    userID,
// 		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
// 	}
// 	header := map[string]string{
// 		"alg": "HS256",
// 		"typ": "JWT",
// 	}
// 	headerBytes, err := json.Marshal(header)
// 	if err != nil {
// 		return "", err
// 	}
// 	payloadBytes, err := json.Marshal(claims)
// 	if err != nil {
// 		return "", err
// 	}
// 	headerEnc := base64urlEncode(headerBytes)
// 	payloadEnc := base64urlEncode(payloadBytes)
// 	message := headerEnc + "." + payloadEnc
// 	h := hmac.New(sha256.New, []byte(secret))
// 	h.Write([]byte(message))
// 	signature := base64urlEncode(h.Sum(nil))
// 	return headerEnc + "." + payloadEnc + "." + signature, nil
// }

// func base64urlEncode(data []byte) string {
// 	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
// }

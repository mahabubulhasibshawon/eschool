package otp

import (
	"context"
	"encoding/json"
	"eschool/database"
	"eschool/util"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/redis/go-redis/v9"
)

type RequestSendOTP struct {
	Username string `json:"username"`
}

// SendOTP generates and stores OTP in Redis, sends via Gmail SMTP
func (h *Handler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var req RequestSendOTP
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("JSON decode error:", err)
		http.Error(w, "Invalid Request Data!", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if user exists
	var user database.User
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE username = $1`
	err = h.middlewares.DB.Get(&user, query, req.Username)
	if err != nil || user.ID == 0 {
		log.Println("User lookup error:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Generate 6-digit OTP
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := fmt.Sprintf("%06d", rng.Intn(1000000))

	// Store OTP in Redis (5-min expiry)
	ctx := context.Background()
	err = h.redisClient.Set(ctx, "otp:"+req.Username, otp, 5*time.Minute).Err()
	if err != nil {
		log.Println("Redis set error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send OTP via Gmail SMTP
	subject := "Subject: Your eSchool OTP\r\n"
	body := fmt.Sprintf("Your one-time password (OTP) is: %s\r\nIt expires in 5 minutes.", otp)
	message := []byte(subject + "\r\n" + body)
	auth := smtp.PlainAuth("", h.smtpUsername, h.smtpPassword, h.smtpHost)
	err = smtp.SendMail(h.smtpHost+":"+h.smtpPort, auth, h.senderEmail, []string{user.Email}, message)
	if err != nil {
		log.Println("SMTP send email error:", err)
		http.Error(w, "Failed to send OTP", http.StatusInternalServerError)
		return
	}

	// Send response
	util.SendData(w, map[string]string{
		"message": "OTP sent to your email",
	}, http.StatusOK)
}

// VerifyOTP checks OTP in Redis
func (h *Handler) VerifyOTP(username, otp string) (bool, error) {
	ctx := context.Background()
	storedOTP, err := h.redisClient.Get(ctx, "otp:"+username).Result()
	if err == redis.Nil || storedOTP != otp {
		log.Println("Redis get error or invalid OTP:", err)
		return false, err
	}
	// Delete OTP after verification
	err = h.redisClient.Del(ctx, "otp:"+username).Err()
	if err != nil {
		log.Println("Redis delete error:", err)
		// Continue, as verification succeeded
	}
	return true, nil
}
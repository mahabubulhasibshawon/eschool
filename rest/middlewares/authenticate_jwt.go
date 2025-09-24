package middlewares

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Claims matches user.go's Claims
type Claims struct {
	UserID    int   `json:"user_id"`
	ExpiresAt int64 `json:"exp"`
}

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 || headerArr[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		// Verify signature
		message := jwtHeader + "." + jwtPayload
		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrMessage := []byte(message)
		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)
		hash := h.Sum(nil)
		newSignature := base64urlEncode(hash)
		if newSignature != signature {
			http.Error(w, "Unauthorized!!!", http.StatusUnauthorized)
			return
		}

		// Decode payload
		payloadBytes, err := base64.RawURLEncoding.DecodeString(jwtPayload)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid payload", http.StatusUnauthorized)
			return
		}
		var claims Claims
		if err := json.Unmarshal(payloadBytes, &claims); err != nil {
			http.Error(w, "Unauthorized: Invalid claims", http.StatusUnauthorized)
			return
		}

		// Check expiration
		if claims.ExpiresAt < time.Now().Unix() {
			http.Error(w, "Unauthorized: Token expired", http.StatusUnauthorized)
			return
		}

		// Set user_id in context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func base64urlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

// package middlewares

// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"net/http"
// 	"strings"
// )

// func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		header := r.Header.Get("Authorization")

// 		if header == "" {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		headerArr := strings.Split(header, " ")

// 		if len(headerArr) != 2 {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		accessToken := headerArr[1]

// 		tokenParts := strings.Split(accessToken, ".")

// 		if len(tokenParts) != 3 {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		}

// 		jwtHeader := tokenParts[0]
// 		jwtPayload := tokenParts[1]
// 		signature := tokenParts[2]

// 		message := jwtHeader + "." + jwtPayload

// 		byteArrSecret := []byte(m.cnf.JwtSecretKey)
// 		byteArrMessage := []byte(message)

// 		h := hmac.New(sha256.New, byteArrSecret)
// 		h.Write(byteArrMessage)

// 		hash := h.Sum(nil)
// 		newSignature := base64urlEncode(hash)

// 		if newSignature != signature {
// 			http.Error(w, "Unauthorized!!!", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// func base64urlEncode(data []byte) string {
// 	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
// }

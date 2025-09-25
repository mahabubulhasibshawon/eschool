package otp

import (
	"context"
	"eschool/rest/middlewares"
	"log"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	middlewares  *middlewares.Middlewares
	redisClient  *redis.Client
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
	senderEmail  string
}

func NewHandler(middlewares *middlewares.Middlewares, redisAddr, redisUsername, redisPassword, smtpHost, smtpPort, smtpUsername, smtpPassword, senderEmail string, redisDB int) *Handler {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: redisUsername,
		Password: redisPassword,
		DB:       redisDB,
		// TLS enabled automatically for non-localhost in go-redis
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	return &Handler{
		middlewares:  middlewares,
		redisClient:  redisClient,
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
		senderEmail:  senderEmail,
	}
}
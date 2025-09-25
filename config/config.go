package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version       string
	ServiceName   string
	HttpPort      int
	JwtSecretKey  string
	JwtRefreshKey string
	RedisAddr     string
	RedisUsrName  string
	RedisPassword string
	RedisDB       int
	SmtpHost      string
	SmtpPort      string
	SmtpUsrName   string
	SmtpPassword  string
	SenderEmail   string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failedto load the env variables: ", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required!")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required!")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required!")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required!")
		os.Exit(1)
	}
	jwtRefreshKey := os.Getenv("JWT_REFRESH_KEY")
	if jwtRefreshKey == "" {
		fmt.Println("Jwt refresh key is required!")
		os.Exit(1)
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		fmt.Println("Redis Address is required!")
		os.Exit(1)
	}
	redisUsrName := os.Getenv("REDIS_USERNAME")
	if redisUsrName == "" {
		fmt.Println("Redis username is required!")
		os.Exit(1)
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		fmt.Println("REDIS PASSWORD is required!")
		os.Exit(1)
	}
	redis_db := os.Getenv("REDIS_DB")
	if redis_db == "" {
		fmt.Println("redis db is required!")
		os.Exit(1)
	}
	db_port, err := strconv.ParseInt(redis_db, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}
	smtpHost := os.Getenv("SMTP_HOST")
	if serviceName == "" {
		fmt.Println("SMTP_HOST is required!")
		os.Exit(1)
	}
	smtpPort := os.Getenv("SMTP_PORT")
	if httpPort == "" {
		fmt.Println("SMTP_PORT is required!")
		os.Exit(1)
	}
	smtpUsername := os.Getenv("SMTP_USERNAME")
	if jwtSecretKey == "" {
		fmt.Println("SMTP_USERNAME is required!")
		os.Exit(1)
	}
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if jwtRefreshKey == "" {
		fmt.Println("SMTP_PASSWORD is required!")
		os.Exit(1)
	}
	senderEmail := os.Getenv("SENDER_EMAIL")
	if redisAddr == "" {
		fmt.Println("SENDER_EMAIL is required!")
		os.Exit(1)
	}

	configurations = &Config{
		Version:       version,
		ServiceName:   serviceName,
		HttpPort:      int(port),
		JwtSecretKey:  jwtSecretKey,
		JwtRefreshKey: jwtRefreshKey,
		RedisAddr:     redisAddr,
		RedisUsrName:  redisUsrName,
		RedisPassword: redisPassword,
		RedisDB:       int(db_port),
		SmtpHost:      smtpHost,
		SmtpPort:      smtpPort,
		SmtpUsrName:   smtpUsername,
		SmtpPassword:  smtpPassword,
		SenderEmail:   senderEmail,
	}
}
func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}

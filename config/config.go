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

	configurations = &Config{
		Version:       version,
		ServiceName:   serviceName,
		HttpPort:      int(port),
		JwtSecretKey:  jwtSecretKey,
		JwtRefreshKey: jwtRefreshKey,
	}
}
func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}

package cmd

import (
	"eschool/config"
	"eschool/rest"
	"eschool/rest/handlers/course"
	"eschool/rest/handlers/otp"
	"eschool/rest/handlers/user"
	"eschool/rest/middlewares"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func Serve() {
	cnf := config.GetConfig()

	// connect db
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	users := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, users, password, dbname,
	)

	DB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("Error connecting database ", err)
		return
	}

	middlewares := middlewares.NewMiddlewares(cnf, DB)
	courseHandler := course.NewHandler(middlewares)
	otpHandler := otp.NewHandler(middlewares, cnf.RedisAddr, cnf.RedisUsrName, cnf.RedisPassword,cnf.SmtpHost, cnf.SmtpPort, cnf.SmtpUsrName, cnf.SmtpPassword, cnf.SenderEmail, cnf.RedisDB)
	userHandler := user.NewHandler(middlewares, otpHandler)

	server := rest.NewServer(
		cnf,
		courseHandler,
		userHandler,
		otpHandler,
	)
	server.Start()
}

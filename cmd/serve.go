package cmd

import (
	"eschool/config"
	"eschool/rest"
	"eschool/rest/handlers/course"
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
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	DB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("Error connecting database ",err)
		return
	}


	middlewares := middlewares.NewMiddlewares(cnf, DB)
	courseHandler := course.NewHandler(middlewares)

	server := rest.NewServer(
		cnf,
		courseHandler,
	)
	server.Start()
}

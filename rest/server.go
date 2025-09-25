package rest

import (
	"eschool/config"
	"eschool/rest/handlers/course"
	"eschool/rest/handlers/otp"
	"eschool/rest/handlers/user"
	"eschool/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf           *config.Config
	courseHandler *course.Handler
	userHandler   *user.Handler
	otpHandler *otp.Handler
}

func NewServer(
	cnf *config.Config,
	courseHandler *course.Handler,
	userHandler *user.Handler,
	otpHandler *otp.Handler,
) *Server {
	return &Server{
		cnf:           cnf,
		courseHandler: courseHandler,
		userHandler:   userHandler,
		otpHandler: otpHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.courseHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.otpHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on port: ", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}

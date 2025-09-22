package rest

import (
	"eschool/config"
	"eschool/rest/handlers/course"
	"eschool/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf           *config.Config
	courseHandler *course.Handler
}

func NewServer(
	cnf *config.Config,
	courseHandler *course.Handler,
) *Server {
	return &Server{
		cnf:           cnf,
		courseHandler: courseHandler,
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

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on port: ", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server",err)
		os.Exit(1)
	}
}

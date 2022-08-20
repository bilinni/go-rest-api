package main

import (
	"example/golang-api/app/internal/config"
	"example/golang-api/app/internal/user"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	cfg *config.Config
}

func main() {
	fmt.Println("create router")
	router := httprouter.New()
	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(server.Serve(listener))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/padwalab/goservices/logger"
	"github.com/padwalab/goservices/router"
)

// PORT on which microservice runs
const PORT = ":8081"

func main() {
	shutdown := make(chan bool)
	router := router.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"})

	go func() {
		log.Fatal(
			http.ListenAndServe(PORT, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
		shutdown <- true
	}()
	logger.Log("INFO", "MicroService started at", PORT)
	<-shutdown
}

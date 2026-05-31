package main

import (
	"fmt" // for printing to the console
	"net/http" // for making HTTP requests
	"github.com/go-chi/chi" // for web development, it is a lightweight router for building Go HTTP services
	"api_go/internal/handlers" 
	log "github.com/sirupsen/logrus" // for logging, it is a structured logger for Go
)

// install external packages using go mod tidy, which will download the packages and add them to the go.mod file

func main() {
	log.SetReportCaller(true) // to log the file and line number of the log statement
	var r *chi.Mux = chi.NewRouter() // create a new router
	handlers.Handler(r) // register the handlers for the routes i.e the functions that will be called when a request is made to a specific route
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r) // start the server on port 8080 and pass the router as the handler for incoming requests
	if err != nil {
		log.Fatal("Error starting server: ", err) // log the error if the server fails to start
	}
}
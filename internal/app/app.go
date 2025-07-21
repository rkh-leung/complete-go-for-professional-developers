package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// app housing data for application and whether it needs a struct

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	// using logger instead of print debugging
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil // nil is valid error type
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) { // r contains data from client
	fmt.Fprintf(w, "Status is available\n")
}

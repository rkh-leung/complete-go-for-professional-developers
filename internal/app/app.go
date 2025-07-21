package app

import (
	"complete-go-for-professional-developers/internal/api"
	"fmt"
	"log"
	"net/http"
	"os"
)

// app housing data for application and whether it needs a struct

type Application struct {
	Logger         *log.Logger
	WorkoutHanlder *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	// using logger instead of print debugging
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Stores go here

	// Handlers go here
	workoutHandler := api.NewWorkoutHandler() // of type *api.WorkoutHandler
	fmt.Printf("The type of workoutHandler is: %T\n", workoutHandler)

	app := &Application{
		Logger:         logger,
		WorkoutHanlder: workoutHandler,
	}

	return app, nil // nil is valid error type
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) { // r contains data from client
	fmt.Fprintf(w, "Status is available\n")
}

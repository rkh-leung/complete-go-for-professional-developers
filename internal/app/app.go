package app

import (
	"complete-go-for-professional-developers/internal/api"
	"complete-go-for-professional-developers/internal/store"
	"complete-go-for-professional-developers/migrations"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

// app housing data for application and whether it needs a struct

type Application struct {
	Logger         *log.Logger
	WorkoutHanlder *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".") // run at base of dir
	if err != nil {
		panic(err)
	}

	// using logger instead of print debugging
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Stores go here
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// Handlers go here
	workoutHandler := api.NewWorkoutHandler(workoutStore) // of type *api.WorkoutHandler

	app := &Application{
		Logger:         logger,
		WorkoutHanlder: workoutHandler,
		DB:             pgDB,
	}

	return app, nil // nil is valid error type
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) { // r contains data from client
	fmt.Fprintf(w, "Status is available\n")
}

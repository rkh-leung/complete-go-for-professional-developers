package routes

import (
	"complete-go-for-professional-developers/internal/app"

	"github.com/go-chi/chi/v5"
)

// Mux is a struct that computes all the handler has all fields to satisfy native standard library HTTP function interface
func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHanlder.HandleGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHanlder.HandleCreateWorkout)
	r.Put("/workouts/{id}", app.WorkoutHanlder.HandleUpdateByID)

	return r
}

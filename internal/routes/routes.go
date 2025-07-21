package routes

import (
	"complete-go-for-professional-developers/internal/app"

	"github.com/go-chi/chi/v5"
)

// Mux is a struct that computes all the handler has all fields to satisfy native standard library HTTP function interface
func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	return r
}

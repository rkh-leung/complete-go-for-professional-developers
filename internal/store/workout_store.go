package store

import "database/sql"

type Workout struct {
	ID              int            `json:"id"` // struct tag
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	DurationMinutes int            `json:"duration_minutes"`
	CaloriesBurned  int            `json:"calories_burned"`
	Entries         []WorkoutEntry `json:"entries"`
}

type WorkoutEntry struct {
	ID              int      `json:"id"`
	ExerciseName    string   `json:"exercise_name"`
	Sets            int      `json:"sets"`
	Reps            *int     `json:"reps"` // constraint from SQL file
	DurationSeconds *int     `json:"duration_seconds"`
	Weight          *float64 `json:"weight"`
	Notes           string   `json:"notes"`
	OrderIndex      int      `json:"order_index"`
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{db: db}
}

// interface a collection of method signatures and if all method signatures are satisfied, then it can be used as a type
type WorkoutStore interface {
	// [methodname](param)(return type)
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkoutByID(id int64) (*Workout, error)
}

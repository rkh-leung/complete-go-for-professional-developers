package api

import (
	"complete-go-for-professional-developers/internal/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// data can also just be functions or different access layers
type WorkoutHandler struct {
	workoutStore store.WorkoutStore // using interface decouples API layer from database layer
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore, // can shorthand to workoutStore
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id") // URLParam returns a string, better with UUID
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return // critical, if it doesn't return it continues
	}

	// base 10, 64 bit int, the return will still be a number due to ParseInt
	// but it affects what base number goes into the path parameter
	// e.g. /workouts/101 => 5 for base 2
	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	workoutById, err := wh.workoutStore.GetWorkoutByID(workoutID)
	fmt.Printf("Here's your workout for id: %d\n", workoutID)
	fmt.Println(workoutById)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	defer r.Body.Close() // Good practice to signal the end of body stream

	// similar to json.Unmarshal() but it takes []bytes, not usable directly on r.Body
	// json.Unmarshal() can only be used when the entire JSON payload is in memory
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Println(err) // temporary
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}

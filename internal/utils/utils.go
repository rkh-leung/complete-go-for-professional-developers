package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]any // empty interface of Go's any type

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func ReadIDParam(r *http.Request) (int64, error) {
	idParam := chi.URLParam(r, "id") // URLParam returns a string, better with UUID
	if idParam == "" {
		return 0, errors.New("invalid id parameter")
	}
	// base 10, 64 bit int, the return will still be a number due to ParseInt
	// but it affects what base number goes into the path parameter
	// e.g. /workouts/101 => 5 for base 2
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter type")
	}
	return id, nil
}

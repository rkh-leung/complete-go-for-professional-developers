package main

import (
	"complete-go-for-professional-developers/internal/app"
	"fmt"
	"net/http"
	"time"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("The app is running!")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{ // server now stores a pointer to http.Server struct
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) { // r contains data from client
	fmt.Fprintf(w, "Status is available\n")

}

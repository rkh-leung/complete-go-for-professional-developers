package main

import (
	"complete-go-for-professional-developers/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("The app is running!")
}

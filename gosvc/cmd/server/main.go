package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/goarch/gosvc/internal/transport/http"
)

// App is ...
type App struct {
}

// Run ...
func (app *App) Run() error {
	fmt.Println("Setting up Application")

	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8000", handler.Router); err != nil {
		fmt.Println("Failed to run server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go rest API")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up Application")
		fmt.Println(err)
	}
}

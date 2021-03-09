package main

import (
	"fmt"
	"net/http"

	"github.com/goarch/gosvc/internal/comment"
	"github.com/goarch/gosvc/internal/database"
	transportHttp "github.com/goarch/gosvc/internal/transport/http"
)

// App is ...
type App struct {
}

// Run ...
func (app *App) Run() error {
	fmt.Println("Setting up Application")

	var err error

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)
	handler := transportHttp.NewHandler(commentService)

	handler.SetupRoutes()

	if err := http.ListenAndServe(":8008", handler.Router); err != nil {
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

package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/goarch/gosvc/internal/comment"
	"github.com/goarch/gosvc/internal/database"
	transportHttp "github.com/goarch/gosvc/internal/transport/http"
	_ "github.com/sirupsen/logrus"
)

// App is ...
type App struct {
	Name    string
	Version string
}

// Run ...
func (app *App) Run() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(
		logrus.Fields{
			"Appname":    app.Name,
			"Appversion": app.Version,
		},
	).Info("Setting up Application")

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
		logrus.Error("Failed to run server")
		return err
	}

	return nil
}

func main() {

	fmt.Println("Go rest API")
	app := App{
		Name:    "Commenting service",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		logrus.Error("Error starting up Application")
		logrus.Fatal(err)
	}
}

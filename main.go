package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/totoro081295/daily-report-api/db"
	"github.com/totoro081295/daily-report-api/db/migrations"
)

func main() {
	e := echo.New()

	db.ConnectDB()
	// migration
	migrations.Execute()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, os.Getenv("PORT"))
	})

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}

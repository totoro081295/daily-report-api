package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Hallo world")
	})

	port := ":1323"
	e.Logger.Fatal(e.Start(port))
}

package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/totoro081295/daily-report-api/db"
	"github.com/totoro081295/daily-report-api/db/migrations"
	mid "github.com/totoro081295/daily-report-api/middleware"
	"github.com/totoro081295/daily-report-api/token"
)

var database *gorm.DB
var tokenHandler token.Handler
var jwt mid.JWTMiddleware

func setup(e *echo.Echo) {
	database = db.ConnectDB()
	// migration
	migrations.Execute()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))

	var err error
	// mount token handler
	tokenHandler = token.NewTokenHandler()

	// mount jwt middleware
	jwt, err = mid.NewJWTMiddleware(tokenHandler)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
		os.Exit(1)
	}

}

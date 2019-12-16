package main

import (
	"os"

	"github.com/labstack/echo"
	accountR "github.com/totoro081295/daily-report-api/account/repository"
	authC "github.com/totoro081295/daily-report-api/auth/controller"
	authU "github.com/totoro081295/daily-report-api/auth/usecase"
	rTokenR "github.com/totoro081295/daily-report-api/refreshtoken/repository"
)

func main() {
	e := echo.New()
	setup(e)
	accountRepo := accountR.NewAccountRepository(database)
	rTokenRepo := rTokenR.NewRefreshTokenRepository(database)

	authUcase := authU.NewAuthUsecase(accountRepo, rTokenRepo, tokenHandler)

	authC.NewAuthController(e, authUcase)

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}

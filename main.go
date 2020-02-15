package main

import (
	"os"

	"github.com/labstack/echo"
	accountC "github.com/totoro081295/daily-report-api/account/controller"
	accountR "github.com/totoro081295/daily-report-api/account/repository"
	accountU "github.com/totoro081295/daily-report-api/account/usecase"
	authC "github.com/totoro081295/daily-report-api/auth/controller"
	authU "github.com/totoro081295/daily-report-api/auth/usecase"
	projectC "github.com/totoro081295/daily-report-api/project/controller"
	projectR "github.com/totoro081295/daily-report-api/project/repository"
	projectU "github.com/totoro081295/daily-report-api/project/usecase"
	rTokenR "github.com/totoro081295/daily-report-api/refreshtoken/repository"
)

func main() {
	e := echo.New()
	setup(e)
	accountRepo := accountR.NewAccountRepository(database)
	rTokenRepo := rTokenR.NewRefreshTokenRepository(database)
	projectRepo := projectR.NewProjectRepository(database)

	accountUcase := accountU.NewAccountUsecase(accountRepo)
	authUcase := authU.NewAuthUsecase(accountRepo, rTokenRepo, tokenHandler)
	projectUcase := projectU.NewProjectUsecase(projectRepo)

	accountC.NewAccountController(e, accountUcase, tokenHandler, jwt)
	authC.NewAuthController(e, authUcase)
	projectC.NewProjectController(e, projectUcase)

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}

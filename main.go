package main

import (
	"os"

	"github.com/labstack/echo"
	accountC "github.com/totoro081295/daily-report-api/account/controller"
	accountR "github.com/totoro081295/daily-report-api/account/repository"
	accountU "github.com/totoro081295/daily-report-api/account/usecase"
	authC "github.com/totoro081295/daily-report-api/auth/controller"
	authU "github.com/totoro081295/daily-report-api/auth/usecase"
	categoryC "github.com/totoro081295/daily-report-api/category/controller"
	categoryR "github.com/totoro081295/daily-report-api/category/repository"
	categoryU "github.com/totoro081295/daily-report-api/category/usecase"
	dContentC "github.com/totoro081295/daily-report-api/dailycontent/controller"
	dContentR "github.com/totoro081295/daily-report-api/dailycontent/repository"
	dContentU "github.com/totoro081295/daily-report-api/dailycontent/usecase"
	projectC "github.com/totoro081295/daily-report-api/project/controller"
	projectR "github.com/totoro081295/daily-report-api/project/repository"
	projectU "github.com/totoro081295/daily-report-api/project/usecase"
	rTokenR "github.com/totoro081295/daily-report-api/refreshtoken/repository"
)

func main() {
	e := echo.New()
	setup(e)
	accountRepo := accountR.NewAccountRepository(database)
	categoryRepo := categoryR.NewCategoryRepository(database)
	dContentRepo := dContentR.NewDailyContentRepository(database)
	rTokenRepo := rTokenR.NewRefreshTokenRepository(database)
	projectRepo := projectR.NewProjectRepository(database)

	accountUcase := accountU.NewAccountUsecase(accountRepo)
	authUcase := authU.NewAuthUsecase(accountRepo, rTokenRepo, tokenHandler)
	categoryUcase := categoryU.NewCategoryUsecase(categoryRepo)
	dContentUcase := dContentU.NewDailyContentUsecase(dContentRepo)
	projectUcase := projectU.NewProjectUsecase(projectRepo)

	accountC.NewAccountController(e, accountUcase, tokenHandler, jwt)
	authC.NewAuthController(e, authUcase)
	categoryC.NewCategoryController(e, categoryUcase, jwt)
	dContentC.NewDailyContentController(e, dContentUcase, tokenHandler, jwt)
	projectC.NewProjectController(e, projectUcase)

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}

package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/totoro081295/daily-report-api/account/usecase"
	"github.com/totoro081295/daily-report-api/middleware"
	"github.com/totoro081295/daily-report-api/status"
	"github.com/totoro081295/daily-report-api/token"
)

// AccountController account controller
type AccountController struct {
	usecase usecase.AccountUsecase
	token   token.Handler
}

// NewAccountController mount account controller
func NewAccountController(e *echo.Echo, account usecase.AccountUsecase, token token.Handler, jwt middleware.JWTMiddleware) {
	handler := &AccountController{
		usecase: account,
		token:   token,
	}
	e.GET("/accounts", handler.Get, jwt.JWT())
}

// Get アカウント取得
func (c *AccountController) Get(ctx echo.Context) error {
	accountID, err := c.token.GetToken(ctx)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	res, err := c.usecase.Get(accountID)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

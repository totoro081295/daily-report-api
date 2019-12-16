package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/totoro081295/daily-report-api/auth"
	"github.com/totoro081295/daily-report-api/auth/usecase"
	"github.com/totoro081295/daily-report-api/status"
)

// AuthController auth controller
type AuthController struct {
	AuthUsecase usecase.AuthUsecase
}

// NewAuthController mount auth controller
func NewAuthController(e *echo.Echo, us usecase.AuthUsecase) {
	handler := &AuthController{
		AuthUsecase: us,
	}
	e.POST("/auth/login", handler.Login)
}

// Login login
func (c *AuthController) Login(ctx echo.Context) error {
	request := auth.Login{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	// validation
	err = auth.LoginValidate(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := c.AuthUsecase.Login(&request)
	if err != nil {
		return status.ResponseError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, token)
}

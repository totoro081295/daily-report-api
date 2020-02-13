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
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)
	e.POST("/refresh", handler.Refresh)
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

// Logout logout
func (c *AuthController) Logout(ctx echo.Context) error {
	request := auth.Logout{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// validation
	err = auth.LogoutValidate(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.AuthUsecase.Logout(request.ID)
	if err != nil {
		return status.ResponseError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// Refresh refresh
func (c *AuthController) Refresh(ctx echo.Context) error {
	request := auth.Refresh{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	token, err := c.AuthUsecase.Refresh(&request)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, token)
}

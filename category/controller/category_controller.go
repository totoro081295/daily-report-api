package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/totoro081295/daily-report-api/category/usecase"
	"github.com/totoro081295/daily-report-api/middleware"
	"github.com/totoro081295/daily-report-api/status"
)

// CategoryController category controller
type CategoryController struct {
	categoryUcase usecase.CategoryUsecase
}

// NewCategoryController mount category controller
func NewCategoryController(e *echo.Echo, category usecase.CategoryUsecase, jwt middleware.JWTMiddleware) {
	handler := &CategoryController{
		categoryUcase: category,
	}
	e.GET("/categories", handler.List, jwt.JWT())
}

// List category一覧取得
func (c *CategoryController) List(ctx echo.Context) error {
	res, err := c.categoryUcase.List()
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

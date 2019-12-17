package controller

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"github.com/totoro081295/daily-report-api/project/usecase"
	"github.com/totoro081295/daily-report-api/status"
)

// ProjectController project controller
type ProjectController struct {
	projectUsecase usecase.ProjectUsecase
}

// NewProjectController mount project controller
func NewProjectController(e *echo.Echo, project usecase.ProjectUsecase) {
	handler := &ProjectController{
		projectUsecase: project,
	}
	e.GET("/projects/:id", handler.Get)
	e.GET("/projects", handler.List)
}

// Get get a project
func (c *ProjectController) Get(ctx echo.Context) error {
	id := uuid.FromStringOrNil(ctx.Param("id"))
	res, err := c.projectUsecase.Get(id)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

// List get projects
func (c *ProjectController) List(ctx echo.Context) error {
	res, err := c.projectUsecase.List()
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

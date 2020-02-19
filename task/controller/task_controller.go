package controller

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"github.com/totoro081295/daily-report-api/middleware"
	"github.com/totoro081295/daily-report-api/status"
	"github.com/totoro081295/daily-report-api/task/usecase"
)

// TaskController task controller
type TaskController struct {
	taskUcase usecase.TaskUsecase
}

// NewTaskController mount task controller
func NewTaskController(e *echo.Echo, task usecase.TaskUsecase, jwt middleware.JWTMiddleware) {
	handler := &TaskController{
		taskUcase: task,
	}
	e.GET("/tasks", handler.List, jwt.JWT())

}

// List list of my task or all task
func (c *TaskController) List(ctx echo.Context) error {
	idStr := ctx.QueryParam("id")
	accountID := uuid.FromStringOrNil(idStr)
	res, err := c.taskUcase.List(accountID)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

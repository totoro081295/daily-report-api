package controller

import (
	"net/http"
	"time"

	"github.com/totoro081295/daily-report-api/dailycontent"
	"github.com/totoro081295/daily-report-api/token"

	"github.com/labstack/echo"
	"github.com/totoro081295/daily-report-api/dailycontent/usecase"
	"github.com/totoro081295/daily-report-api/middleware"
	"github.com/totoro081295/daily-report-api/status"
)

// DailyContentController dailyContent controller
type DailyContentController struct {
	usecase usecase.DailyContentUsecase
	token   token.Handler
}

// NewDailyContentController mount dailyContent controller
func NewDailyContentController(
	e *echo.Echo,
	dailyContent usecase.DailyContentUsecase,
	token token.Handler,
	jwt middleware.JWTMiddleware,
) {
	handler := &DailyContentController{
		usecase: dailyContent,
		token:   token,
	}

	e.GET("/daily-contents/:date", handler.GetByTargetDate, jwt.JWT())
	e.POST("/daily-contents", handler.Create, jwt.JWT())
	e.PATCH("/daily-contents", handler.Update, jwt.JWT())
}

// GetByTargetDate 対象日のdailyContentを取得する
func (c *DailyContentController) GetByTargetDate(ctx echo.Context) error {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	date, err := time.ParseInLocation("2006-01-02", ctx.Param("date"), loc)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	res, err := c.usecase.GetByTargetDate(date)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

// Create dailyContentを作成する
func (c *DailyContentController) Create(ctx echo.Context) error {
	request := dailycontent.CreatePayload{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	accountID, err := c.token.GetToken(ctx)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	request.CreatedBy = accountID
	res, err := c.usecase.Create(&request)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.JSON(http.StatusCreated, res)
}

// Update dailyContentを更新する
func (c *DailyContentController) Update(ctx echo.Context) error {
	request := dailycontent.UpdatePayload{}
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	accountID, err := c.token.GetToken(ctx)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	request.UpdatedBy = accountID
	err = c.usecase.Update(request)
	if err != nil {
		return status.ResponseError(ctx, err)
	}
	return ctx.NoContent(http.StatusNoContent)

}

package status

import (
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/labstack/echo"
)

var (
	// ErrInternalServer internal server error
	ErrInternalServer = errors.New("Internal server error")
	// ErrNotFound not found error
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict conflit error
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadRequest bad request error
	ErrBadRequest = errors.New("Bad request")
	// ErrUnauthrized unauthorized error
	ErrUnauthrized = errors.New("Unauthorized")
	// ErrForbidden forbidden error
	ErrForbidden = errors.New("Forbidden")
)

// ErrorMessage error message
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResponseError 返却するエラーコードの指定
func ResponseError(ctx echo.Context, err error) error {
	res := ErrorMessage{}
	if grpc.Code(err) == codes.NotFound {
		var errMsg = ErrorMessage{
			Code:    404,
			Message: "Resource with ID:" + ctx.Param("id") + " is not found.",
		}
		return ctx.JSON(404, errMsg)
	}
	switch errors.Cause(err) {
	case ErrInternalServer:
		res.Code = 500
		res.Message = err.Error()
		return ctx.JSON(500, res)
	case ErrNotFound:
		res.Code = 404
		res.Message = err.Error()
		return ctx.JSON(http.StatusNotFound, res)
	case ErrConflict:
		res.Code = http.StatusConflict
		res.Message = err.Error()
		return ctx.JSON(http.StatusConflict, res)
	case ErrBadRequest:
		res.Code = 400
		res.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, res)
	case ErrUnauthrized:
		res.Code = http.StatusUnauthorized
		res.Message = err.Error()
		return ctx.JSON(http.StatusUnauthorized, res)
	case ErrForbidden:
		res.Code = http.StatusForbidden
		res.Message = err.Error()
		return ctx.JSON(http.StatusForbidden, res)
	default:
		res.Code = 500
		res.Message = err.Error()
		return ctx.JSON(500, res)
	}
}

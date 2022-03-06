package delivery

import (
	"fmt"
	"net/http"

	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorHandler :nodoc:
func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed || err == nil {
		return
	}

	if c.Response().Size > 0 {
		return
	}

	httpStatusCode := http.StatusInternalServerError
	resp := ErrResponse{
		Code:    "-1",
		Message: "fatal error! please contact the service owner",
	}

	switch e := err.(type) {
	case *internal.Error:
		httpStatusCode = http.StatusBadRequest
		resp.Message = e.Error()
		resp.Code = e.StatusCode()
	case *echo.HTTPError:
		httpStatusCode = e.Code
		resp.Message = e.Error()
		resp.Code = cast.ToString(e.Code)
	default:
		fmt.Println("FATAL ERROR:", e.Error())
	}

	_ = c.JSON(httpStatusCode, resp)
}

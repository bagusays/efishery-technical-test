package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *server) FetchAllResource(c echo.Context) error {
	result, err := h.resourceService.FetchResource(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

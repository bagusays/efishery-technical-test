package delivery

import (
	"net/http"
	"strings"

	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/labstack/echo/v4"
)

func (h *server) ValidateToken(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	token = strings.Replace(token, "Bearer ", "", -1)
	if token == "" {
		return internal.ErrTokenIsMissing
	}

	userClaim, err := h.authService.ValidateToken(token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userClaim)
}

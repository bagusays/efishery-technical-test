package delivery

import (
	"strings"

	"github.com/bagusays/efishery-technical-test/internal"
	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// AuthorizeFor :nodoc:
func AuthorizeFor(roles ...model.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, err := extractRoleFromToken(c)
			if err != nil {
				return err
			}

			for _, d := range roles {
				if role == d {
					return next(c)
				}
			}

			return internal.ErrUnauthorized
		}
	}
}

func extractRoleFromToken(c echo.Context) (role model.Role, err error) {
	tokenString := c.Request().Header.Get(echo.HeaderAuthorization)
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	if tokenString == "" {
		return "", internal.ErrTokenIsMissing
	}

	token, err := jwt.ParseWithClaims(tokenString, &model.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtSecret), nil
	})

	if err != nil {
		return "", internal.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*model.UserClaim); ok && token.Valid {
		return claims.Role, nil
	}

	return "", internal.ErrInvalidToken
}

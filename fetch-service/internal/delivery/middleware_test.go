package delivery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/bagusays/efishery-technical-test/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizeFor(t *testing.T) {
	testCases := []struct {
		name          string
		userRole      model.Role
		authorizeRole []model.Role
		isErr         bool
	}{
		{
			name:          "success with admin only",
			userRole:      model.RoleAdmin,
			authorizeRole: []model.Role{model.RoleAdmin},
		},
		{
			name:          "success with basic and admin",
			userRole:      model.RoleBasic,
			authorizeRole: []model.Role{model.RoleBasic, model.RoleAdmin},
		},
		{
			name:          "failed with admin only",
			userRole:      model.RoleBasic,
			authorizeRole: []model.Role{model.RoleAdmin},
			isErr:         true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			server := echo.New()

			next := func(c echo.Context) error {
				return nil
			}

			middleware := Middleware{
				jwtSecret: "123",
			}

			h := middleware.AuthorizeFor(tt.authorizeRole...)(next)

			tokenString, err := service.GenerateJWT("123", model.User{
				Role: tt.userRole,
			})
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/ini-url", nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenString)

			rec := httptest.NewRecorder()
			c := server.NewContext(req, rec)

			err = h(c)
			if tt.isErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}

}

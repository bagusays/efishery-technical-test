package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		auth := &auth{
			jwtSecret: "yCsWIzH67sxldDC5K9l55ZxKYjiFoUjugmYLz3Nv",
		}

		userClaim := &model.UserClaim{
			User: model.User{
				Phone:      "1",
				Name:       "2",
				Role:       "3",
				UserName:   "4",
				Created_at: time.Time{},
			},
		}
		tokenString, err := GenerateJWT(auth.jwtSecret, userClaim.User)
		fmt.Println(tokenString)
		assert.NoError(t, err)
		result, err := auth.ValidateToken(tokenString)
		assert.NoError(t, err)
		assert.EqualValues(t, userClaim.User, result.User)
	})

	t.Run("invalid token", func(t *testing.T) {
		auth := &auth{
			jwtSecret: "test",
		}
		_, err := auth.ValidateToken("test")
		assert.Error(t, err)
	})
}

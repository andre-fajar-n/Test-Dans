package middleware

import (
	"dans/pkg"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Auth struct {
	secretKey string
}

func NewAuth(cfg *viper.Viper) Auth {
	return Auth{
		secretKey: cfg.Sub("app").GetString("secret_key"),
	}
}

func (a *Auth) RequiredToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header

		authHeader := header.Get("authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, errors.New("token is empty"))
		}

		sp := strings.Split(authHeader, " ")
		var jwtString string

		if len(sp) == 1 {
			jwtString = sp[0] // Authorization: yourJWThere
		} else {
			jwtString = sp[1] // Authorization: Bearer yourJWThere
		}

		_, err := pkg.ValidateToken(jwtString, a.secretKey)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		return next(c)
	}
}

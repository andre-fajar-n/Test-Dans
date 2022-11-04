package middleware

import (
	"dans/pkg"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
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

func (a *Auth) RequiredToken() echo.MiddlewareFunc {
	configMiddleware := echomiddleware.JWTConfig{
		Claims:     &pkg.Claims{},
		SigningKey: []byte(a.secretKey),
	}

	return echomiddleware.JWTWithConfig(configMiddleware)
}

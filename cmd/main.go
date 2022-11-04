package main

import (
	"dans/env"
	"dans/handler"
	"dans/pkg"
	"dans/postgre"
	"dans/thirdparty/dans"
	"dans/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	cfg := env.InitConfig()

	db := postgre.NewConnection(cfg)

	postgre.Migrate(db)

	// Postgre
	userPostgre := postgre.NewUserPostgre(db)

	// Third Party
	// Dans
	dansJobApi := dans.NewJob(cfg)

	// Usecase
	userUsecase := usecase.NewUser(userPostgre, cfg)
	jobUsecase := usecase.NewJob(dansJobApi)

	// Handler
	userHandler := handler.NewUser(userUsecase)
	jobHandler := handler.NewJob(jobUsecase)

	// Middleware
	configMiddleware := echomiddleware.JWTConfig{
		Claims:     &pkg.Claims{},
		SigningKey: []byte(cfg.Sub("app").GetString("secret_key")),
	}

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetLevel(log.INFO)

	// v1
	v1 := e.Group("/v1")
	{
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		v1Job := v1.Group("/job", echomiddleware.JWTWithConfig(configMiddleware))
		v1Job.GET("", jobHandler.GetList)
		v1Job.GET("/:id", jobHandler.GetDetail)
	}

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

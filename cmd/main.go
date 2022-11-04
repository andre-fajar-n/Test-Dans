package main

import (
	"dans/env"
	"dans/handler"
	"dans/postgre"
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

	// Usecase
	userUsecase := usecase.NewUser(userPostgre, cfg)

	// Handler
	userHandler := handler.NewUser(userUsecase)

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetLevel(log.INFO)

	// v1
	v1 := e.Group("/v1")
	{
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)
	}

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

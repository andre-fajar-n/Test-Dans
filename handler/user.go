package handler

import (
	"dans/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	userUsecase entity.UserUsecase
}

func NewUser(
	userUsecase entity.UserUsecase,
) User {
	return User{
		userUsecase,
	}
}

func (h *User) Register(c echo.Context) error {
	req := entity.UserRegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.userUsecase.Register(c.Request().Context(), &req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success registered")
}

func (h *User) Login(c echo.Context) error {
	req := entity.UserLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	output, err := h.userUsecase.Login(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, output)
}

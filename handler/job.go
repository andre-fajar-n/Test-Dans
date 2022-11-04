package handler

import (
	"dans/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Job struct {
	userUsecase entity.JobUsecase
}

func NewJob(
	userUsecase entity.JobUsecase,
) Job {
	return Job{
		userUsecase,
	}
}

func (h *Job) GetDetail(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusUnprocessableEntity, "id cannot be empty")
	}

	result, err := h.userUsecase.GetDetail(c.Request().Context(), id)
	if err != nil {
		fmt.Println("HANDLER", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

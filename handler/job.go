package handler

import (
	"dans/entity"
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Job) GetList(c echo.Context) error {
	req := entity.JobListRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.userUsecase.GetList(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

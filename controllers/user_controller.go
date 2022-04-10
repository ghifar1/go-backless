package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func (res UserController) Get(c echo.Context) error {
	res.Status = http.StatusOK
	res.Data = "Welcome to Backless!"
	return c.JSON(http.StatusOK, res)
}

package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WelcomeController struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func (res WelcomeController) Get(c echo.Context) error {
	res.Status = http.StatusOK
	res.Data = "Welcome to Backless!"
	return c.JSON(http.StatusOK, res)

}

func (res WelcomeController) Post(c echo.Context) error {
	res.Status = http.StatusOK
	res.Data = "Welcome to Backless!"
	return c.JSON(http.StatusOK, res)
}

package routers

import (
	"backless/controllers"
	"os"

	"github.com/labstack/echo/v4"
)

func Api() {
	e := echo.New()

	e.GET("/", controllers.WelcomeController{}.Get)
	e.POST("/", controllers.WelcomeController{}.Post)

	e.GET("/user", controllers.UserController{}.Get)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}

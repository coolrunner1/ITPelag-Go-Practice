package router

import (
	"github.com/coolrunner1/project/controllers"
	"github.com/labstack/echo/v4"
)

func TestRoutes(app *echo.Echo) {
	group := app.Group("/")
	group.GET("", controllers.GetTestData)
}

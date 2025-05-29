package router

import (
	"github.com/coolrunner1/project/cmd/controllers"
	"github.com/labstack/echo/v4"
)

func CommentRoutes(app *echo.Echo) {
	group := app.Group("/api/v1/comments")
	group.GET("", controllers.GetComments)
}

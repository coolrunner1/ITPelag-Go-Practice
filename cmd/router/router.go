package router

import (
	"github.com/coolrunner1/project/cmd/handler"
	"github.com/labstack/echo/v4"
)

func commentRoutes(app *echo.Echo) {
	group := app.Group("/api/v1/comments")
	group.GET("", handler.GetComments)
}

func categoryRoutes(app *echo.Echo) {
	group := app.Group("/api/v1/categories")
	group.GET("", handler.GetCategories)
	group.GET("/:id", handler.GetCategory)
	group.POST("", handler.PostCategory)
	group.PUT("/:id", handler.PutCategory)
	group.DELETE("/:id", handler.DeleteCategory)
}

func GetRoutes(app *echo.Echo) {
	commentRoutes(app)
	categoryRoutes(app)
}

package router

import (
	"github.com/coolrunner1/project/cmd/controllers"
	"github.com/labstack/echo/v4"
)

func commentRoutes(app *echo.Echo) {
	group := app.Group("/api/v1/comments")
	group.GET("", controllers.GetComments)
}

func categoryRoutes(app *echo.Echo) {
	group := app.Group("/api/v1/categories")
	group.GET("", controllers.GetCategories)
	group.GET("/:id", controllers.GetCategory)
	group.POST("", controllers.PostCategory)
	group.PUT("/:id", controllers.PutCategory)
	group.DELETE("/:id", controllers.DeleteCategory)
}

func GetRoutes(app *echo.Echo) {
	commentRoutes(app)
	categoryRoutes(app)
}

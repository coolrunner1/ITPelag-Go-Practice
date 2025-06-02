package router

import (
	"github.com/coolrunner1/project/cmd/handler"
	"github.com/coolrunner1/project/cmd/repository"
	"github.com/coolrunner1/project/cmd/storage"
	"github.com/labstack/echo/v4"
)

func commentRoutes(app *echo.Echo) {
	group := app.Group("/api/v1/comments")
	group.GET("", handler.GetComments)
}

func categoryRoutes(app *echo.Echo) {
	db := storage.GetDB()
	if db == nil {
		panic("DB Not Found")
	}
	categoryHandler := handler.NewCategoryHandler(repository.NewCategoryRepository(db))
	group := app.Group("/api/v1/categories")
	group.GET("", categoryHandler.GetCategories)
	group.GET("/:id", categoryHandler.GetCategory)
	group.POST("", categoryHandler.PostCategory)
	group.PUT("/:id", categoryHandler.PutCategory)
	group.DELETE("/:id", categoryHandler.DeleteCategory)
}

func GetRoutes(app *echo.Echo) {
	commentRoutes(app)
	categoryRoutes(app)
}

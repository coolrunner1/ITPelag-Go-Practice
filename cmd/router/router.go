package router

import (
	"github.com/coolrunner1/project/cmd/handler"
	"github.com/coolrunner1/project/cmd/repository"
	"github.com/coolrunner1/project/cmd/service"
	"github.com/coolrunner1/project/cmd/storage"
	"github.com/go-playground/validator/v10"
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

	categoryHandler := handler.NewCategoryHandler(service.NewCategoryService(repository.NewCategoryRepository(db), validator.New()))
	group := app.Group("/api/v1/categories")
	group.GET("", categoryHandler.GetCategories)
	group.GET("/:id", categoryHandler.GetCategory)
	group.POST("", categoryHandler.PostCategory)
	group.PUT("/:id", categoryHandler.PutCategory)
	group.DELETE("/:id", categoryHandler.DeleteCategory)
}

func userRoutes(app *echo.Echo) {
	db := storage.GetDB()
	if db == nil {
		panic("DB Not Found")
	}
	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db)))
	group := app.Group("/api/v1/users")
	group.GET("", userHandler.GetUsers)
	group.GET("/:id", userHandler.GetUserById)
}

func authRoutes(app *echo.Echo) {
	db := storage.GetDB()
	if db == nil {
		panic("DB Not Found")
	}
	authHandler := handler.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db), validator.New()))
	group := app.Group("/api/v1/auth")
	group.POST("/register", authHandler.Register)
	group.POST("/login", authHandler.Login)
}

func GetRoutes(app *echo.Echo) {
	commentRoutes(app)
	categoryRoutes(app)
	userRoutes(app)
	authRoutes(app)
}

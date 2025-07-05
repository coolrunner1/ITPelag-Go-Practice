package router

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/handler"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/coolrunner1/project/internal/service"
	"github.com/coolrunner1/project/internal/storage"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Router interface {
	GetRoutes()
}

type router struct {
	app *echo.Echo
	db  *sql.DB
}

func NewRouter(app *echo.Echo) Router {
	db := storage.GetDB()

	if db == nil {
		panic("DB Not Found")
	}
	return &router{
		app: app,
		db:  db,
	}
}

func (r *router) commentRoutes() {
	group := r.app.Group("/api/v1/comments")
	group.GET("", handler.GetComments)
}

func (r *router) categoryRoutes() {
	categoryHandler := handler.NewCategoryHandler(service.NewCategoryService(repository.NewCategoryRepository(r.db), validator.New()))
	group := r.app.Group("/api/v1/categories")
	group.GET("", categoryHandler.GetCategories)
	group.GET("/:id", categoryHandler.GetCategory)
	group.POST("", categoryHandler.PostCategory)
	group.PUT("/:id", categoryHandler.PutCategory)
	group.DELETE("/:id", categoryHandler.DeleteCategory)
}

func (r *router) userRoutes() {
	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(r.db), validator.New()))
	group := r.app.Group("/api/v1/users")
	group.GET("", userHandler.GetUsers)
	group.GET("/:id", userHandler.GetUserById)
	group.PUT("/:id", userHandler.UpdateUser)
	group.DELETE("/:id", userHandler.DeleteUser)
}

func (r *router) authRoutes() {
	authHandler := handler.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(r.db), validator.New()))
	group := r.app.Group("/api/v1/auth")
	group.POST("/register", authHandler.Register)
	group.POST("/login", authHandler.Login)
}

func (r *router) searchRoutes() {
	searchHandler := handler.NewSearchHandler()
	group := r.app.Group("/api/v1/search")
	group.GET("/posts", searchHandler.SearchPosts)
	group.GET("/users", searchHandler.SearchUsers)
	group.GET("/communities", searchHandler.SearchCommunities)
}

func (r *router) postRoutes() {
	postsHandler := handler.NewPostsHandler()
	group := r.app.Group("/api/v1")
	group.GET("/posts", postsHandler.GetAllPosts)
	group.GET("/posts/:id", postsHandler.GetPostById)
	group.POST("/posts", postsHandler.CreatePost)
	group.PUT("/posts/:id", postsHandler.UpdatePost)
	group.DELETE("/posts/:id", postsHandler.DeletePost)
	group.GET("/communities/:id/posts", postsHandler.GetPostsByCommunity)
	group.GET("/users/:id/posts", postsHandler.GetPostsByUser)
}

func (r *router) communityRoutes() {
	communityHandler := handler.NewCommunityHandler(
		service.NewCommunityService(
			repository.NewCommunityRepository(
				r.db,
				repository.NewCategoryRepository(r.db),
				repository.NewUserRepository(r.db),
			),
			validator.New(),
		),
	)
	group := r.app.Group("/api/v1/communities")
	group.GET("", communityHandler.GetAllCommunities)
}

func (r *router) GetRoutes() {
	r.commentRoutes()
	r.categoryRoutes()
	r.userRoutes()
	r.authRoutes()
	r.searchRoutes()
	r.postRoutes()
	r.communityRoutes()
}

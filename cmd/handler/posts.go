package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PostsHandler interface {
	GetAllPosts(c echo.Context) error
	GetPostById(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
	GetPostsByCommunity(c echo.Context) error
	GetPostsByUser(c echo.Context) error
}

type postsHandler struct {
}

func NewPostsHandler() PostsHandler {
	return &postsHandler{}
}

func (p *postsHandler) GetAllPosts(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (p *postsHandler) GetPostById(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (p *postsHandler) CreatePost(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}
func (p *postsHandler) UpdatePost(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (p *postsHandler) DeletePost(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (p *postsHandler) GetPostsByCommunity(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (p *postsHandler) GetPostsByUser(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

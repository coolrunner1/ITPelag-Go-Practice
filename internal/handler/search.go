package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SearchHandler interface {
	SearchPosts(c echo.Context) error
	SearchUsers(c echo.Context) error
	SearchCommunities(c echo.Context) error
}

type searchHandler struct {
}

func NewSearchHandler() SearchHandler {
	return &searchHandler{}
}

func (h *searchHandler) SearchPosts(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "SearchPosts not implemented")
}

func (h *searchHandler) SearchUsers(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "SearchUsers not implemented")
}

func (h *searchHandler) SearchCommunities(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "SearchCommunities not implemented")
}

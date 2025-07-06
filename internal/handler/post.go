package handler

import (
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/service"
	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	postService service.PostService
}

func NewPostsHandler(postService service.PostService) PostsHandler {
	return &postsHandler{
		postService: postService,
	}
}

func (h *postsHandler) GetAllPosts(c echo.Context) error {
	start, err := strconv.Atoi(c.QueryParam("start"))
	if err != nil {
		start = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 15
	}

	posts, err := h.postService.GetAll(start, limit)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *postsHandler) GetPostById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
	}
	post, err := h.postService.GetById(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (h *postsHandler) CreatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid community id")
	}

	req := &dto.PostCreateRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.postService.Create(req, 1, id)
	if err != nil {
		if errors.Is(err, service.ErrValidation) || errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *postsHandler) UpdatePost(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *postsHandler) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	err = h.postService.DeleteById(id, 1)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "No post found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *postsHandler) GetPostsByCommunity(c echo.Context) error {
	start, err := strconv.Atoi(c.QueryParam("start"))
	if err != nil {
		start = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 15
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid community id")
	}

	posts, err := h.postService.GetAllByCommunityId(start, limit, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *postsHandler) GetPostsByUser(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

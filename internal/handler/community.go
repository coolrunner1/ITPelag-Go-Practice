package handler

import (
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/service"
	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CommunityHandler interface {
	GetAllCommunities(c echo.Context) error
	CreateCommunity(c echo.Context) error
	GetCommunityById(c echo.Context) error
	UpdateCommunity(c echo.Context) error
	DeleteCommunity(c echo.Context) error
	GetCommunityByUserSubscriptions(c echo.Context) error
	SubscribeToCommunity(c echo.Context) error
	UnsubscribeFromCommunity(c echo.Context) error
	GetCommunitySubscribers(c echo.Context) error
	GetCommunityModerators(c echo.Context) error
}

type communityHandler struct {
	communityService service.CommunityService
}

func NewCommunityHandler(communityService service.CommunityService) CommunityHandler {
	return &communityHandler{
		communityService: communityService,
	}
}

func (h *communityHandler) GetAllCommunities(c echo.Context) error {
	start, err := strconv.Atoi(c.QueryParam("start"))
	if err != nil {
		start = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 15
	}

	communities, err := h.communityService.GetAll(start, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, communities)
}

func (h *communityHandler) GetCommunityById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid community id")
	}
	community, err := h.communityService.GetById(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Community not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, community)
}

func (h *communityHandler) CreateCommunity(c echo.Context) error {
	req := &dto.CommunityCreateRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.communityService.Create(req, 1)
	if err != nil {
		if errors.Is(err, service.ErrValidation) || errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *communityHandler) UpdateCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid community ID")
	}
	req := &dto.CommunityUpdateRequest{}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := h.communityService.Update(req, id, 1)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Community not found")
		}
		if errors.Is(err, service.ErrValidation) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updated)
}

func (h *communityHandler) DeleteCommunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	err = h.communityService.DeleteById(id, 1)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "No community found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// TODO: rename
func (h *communityHandler) GetCommunityByUserSubscriptions(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *communityHandler) SubscribeToCommunity(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *communityHandler) UnsubscribeFromCommunity(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *communityHandler) GetCommunitySubscribers(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *communityHandler) GetCommunityModerators(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

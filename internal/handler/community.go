package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
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
}

type communityHandler struct {
}

func NewCommunityHandler() CommunityHandler {
	return &communityHandler{}
}

func (h *communityHandler) GetAllCommunities(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *communityHandler) GetCommunityById(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *communityHandler) UpdateCommunity(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

func (h *communityHandler) DeleteCommunity(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

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

func (h *communityHandler) CreateCommunity(c echo.Context) error {
	//TODO implement me
	return c.JSON(http.StatusNotImplemented, "Not Implemented")
}

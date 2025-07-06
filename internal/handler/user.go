package handler

import (
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/service"
	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler interface {
	GetUsers(c echo.Context) error
	GetUserById(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetMyUser(c echo.Context) error
	UpdateMyUser(c echo.Context) error
	DeleteMyUser(c echo.Context) error
	RestoreUser(c echo.Context) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	start, err := strconv.Atoi(c.QueryParam("start"))
	if err != nil {
		start = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 15
	}

	users, err := h.userService.GetAll(start, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}
	user, err := h.userService.GetById(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	req := &dto.UserUpdateRequest{}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := h.userService.Update(*req, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		if errors.Is(err, service.ErrValidation) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updated)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	if err := h.userService.DeleteById(id); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		if errors.Is(err, service.ErrAlreadyDeleted) {
			return echo.NewHTTPError(http.StatusConflict, "User already deleted")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *userHandler) RestoreUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	user, err := h.userService.RestoreById(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		if errors.Is(err, service.ErrNotDeleted) {
			return echo.NewHTTPError(http.StatusConflict, "User hasn't been deleted")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetMyUser(c echo.Context) error {
	user, err := h.userService.GetById(1)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateMyUser(c echo.Context) error {
	req := &dto.UserUpdateRequest{}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := h.userService.Update(*req, 1)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		if errors.Is(err, service.ErrValidation) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updated)
}

func (h *userHandler) DeleteMyUser(c echo.Context) error {
	return c.JSON(http.StatusTeapot, "Implemented, but temporarily disabled")
	if err := h.userService.DeleteById(1); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		if errors.Is(err, service.ErrAlreadyDeleted) {
			return echo.NewHTTPError(http.StatusConflict, "User already deleted")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

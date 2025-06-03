package handler

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/service"
	"github.com/go-errors/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler interface {
	GetUsers(c echo.Context) error
	GetUserById(c echo.Context) error
}

type userHandler struct {
	userService service.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
		validator:   validator.New(),
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
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "No users found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUserById(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "Not implemented")
}

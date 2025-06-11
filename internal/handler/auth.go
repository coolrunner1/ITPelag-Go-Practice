package handler

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/dto"
	"github.com/coolrunner1/project/internal/service"
	"github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) Register(c echo.Context) error {
	req := &dto.RegisterRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.authService.Register(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(c echo.Context) error {
	req := &dto.LoginRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res, err := h.authService.Login(*req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
		}
		if errors.Is(err, service.ErrValidation) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

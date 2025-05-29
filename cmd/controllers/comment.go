package controllers

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CommentController interface {
	GetComments(c echo.Context) error
}

func GetComments(c echo.Context) error {
	comments, err := repositories.GetComments()
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No comments found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comments)
}

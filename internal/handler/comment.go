package handler

import (
	"database/sql"
	"github.com/coolrunner1/project/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetComments(c echo.Context) error {
	comments, err := repository.GetComments()
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No comments found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comments)
}

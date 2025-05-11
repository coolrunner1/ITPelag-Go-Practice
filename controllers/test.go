package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type TestController interface {
	GetTestData(c echo.Context) error
}

func GetTestData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"hello": "world",
	})
}

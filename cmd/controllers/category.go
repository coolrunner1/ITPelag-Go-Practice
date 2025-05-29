package controllers

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/models"
	"github.com/coolrunner1/project/cmd/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCategories(c echo.Context) error {
	categories, err := repositories.GetCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No categories found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	id := c.Param("id")
	category, err := repositories.GetCategoryById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func PostCategory(c echo.Context) error {
	category := &models.Category{}
	if err := c.Bind(category); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if category.Title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title is required")
	}
	category, err := repositories.CreateCategory(*category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, category)
}

func PutCategory(c echo.Context) error {
	id := c.Param("id")
	category := &models.Category{}
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if category.Title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title is required")
	}
	if _, err := repositories.GetCategoryById(id); err != nil {
		if err == sql.ErrNoRows {
			category, err := repositories.CreateCategory(*category)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusCreated, category)
		}
	}
	category, err := repositories.UpdateCategory(*category, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	if _, err := repositories.GetCategoryById(id); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category by id "+id+" found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err := repositories.DeleteCategoryById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, "")
}

package handler

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/model"
	"github.com/coolrunner1/project/cmd/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCategories(c echo.Context) error {
	categories, err := repository.GetCategories()
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
	category, err := repository.GetCategoryById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func PostCategory(c echo.Context) error {
	category := &model.Category{}
	if err := c.Bind(category); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if category.Title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title is required")
	}
	category, err := repository.CreateCategory(*category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, category)
}

func PutCategory(c echo.Context) error {
	id := c.Param("id")
	category := &model.Category{}
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if category.Title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title is required")
	}
	if _, err := repository.GetCategoryById(id); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Category not found")
		}
	}
	category, err := repository.UpdateCategory(*category, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	if _, err := repository.GetCategoryById(id); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category by id "+id+" found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err := repository.DeleteCategoryById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, "")
}

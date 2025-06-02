package handler

import (
	"database/sql"
	"github.com/coolrunner1/project/cmd/dto"
	"github.com/coolrunner1/project/cmd/model"
	"github.com/coolrunner1/project/cmd/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CategoryHandler interface {
	GetCategories(c echo.Context) error
	GetCategory(c echo.Context) error
	PostCategory(c echo.Context) error
	PutCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}

type categoryHandler struct {
	categoryRepo repository.CategoryRepository
	validator    *validator.Validate
}

func NewCategoryHandler(categoryRepo repository.CategoryRepository) CategoryHandler {
	return &categoryHandler{
		categoryRepo: categoryRepo,
		validator:    validator.New(),
	}
}

func (ch *categoryHandler) GetCategories(c echo.Context) error {
	categories, err := ch.categoryRepo.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No categories found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, categories)
}

func (ch *categoryHandler) GetCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}
	category, err := ch.categoryRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func (ch *categoryHandler) PostCategory(c echo.Context) error {
	categoryDTO := &dto.CategoryRequest{}
	if err := c.Bind(categoryDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ch.validator.Struct(categoryDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation failed: "+err.Error())
	}

	category := model.Category{
		Title: categoryDTO.Title,
	}

	created, err := ch.categoryRepo.Create(category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, created)
}

func (ch *categoryHandler) PutCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}
	categoryDTO := &dto.CategoryRequest{}

	if err := c.Bind(categoryDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ch.validator.Struct(categoryDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation failed: "+err.Error())
	}

	category := model.Category{
		Title: categoryDTO.Title,
	}
	if _, err := ch.categoryRepo.GetById(id); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Category not found")
		}
	}
	updated, err := ch.categoryRepo.Update(category, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updated)
}

func (ch *categoryHandler) DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}
	if _, err := ch.categoryRepo.GetById(id); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = ch.categoryRepo.DeleteById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "No category found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, "")
}

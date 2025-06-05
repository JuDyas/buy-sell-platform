package handler

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CategoryHandler struct {
	service service.CategoriesService
}

func NewCategoryHandler(service service.CategoriesService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.CategoryCreate
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		id, err := h.service.Create(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot create category"})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
	}
}

func (h *CategoryHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			req dto.CategoryUpdate
			id  = c.Param("id")
		)

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		idObj, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		if err := h.service.Update(c.Request().Context(), idObj, req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot update category"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "category updated"})
	}
}

func (h *CategoryHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idObj, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		category, err := h.service.GetByID(c.Request().Context(), idObj)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
		}

		return c.JSON(http.StatusOK, category)
	}
}

func (h *CategoryHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idObj, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		err = h.service.Delete(c.Request().Context(), idObj)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot delete category"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "category deleted"})
	}
}

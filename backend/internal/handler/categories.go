package handler

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryHandler struct {
	service service.CategoriesService
}

func NewCategoryHandler(service *service.CategoriesService) *CategoryHandler {
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

package handler

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type AdvertHandler struct {
	service service.AdvertService
}

func NewAdvertHandler(service service.AdvertService) *AdvertHandler {
	return &AdvertHandler{
		service: service,
	}
}

func (h *AdvertHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.AdvertCreate
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		authorIDStr := c.Get("userID")
		if authorIDStr == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}

		authorID, err := primitive.ObjectIDFromHex(authorIDStr.(string))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid author id"})
		}

		advert, err := h.service.Create(c.Request().Context(), authorID, req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot create advert"})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{"id": advert.ID.Hex()})
	}
}

func (h *AdvertHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.AdvertUpdate
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		userIDStr := c.Get("userID")
		if userIDStr == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}

		userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		}

		advertID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		advert, err := h.service.GetByID(c.Request().Context(), advertID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "advert not found"})
		}

		if advert.AuthorID != userID {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}

		err = h.service.Update(c.Request().Context(), advertID, req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot update advert"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "advert updated"})
	}
}

func (h *AdvertHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		advert, err := h.service.GetByID(c.Request().Context(), advertID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "advert not found"})
		}

		return c.JSON(http.StatusOK, advert)
	}
}

func (h *AdvertHandler) SoftDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		advertID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		err = h.service.SoftDelete(c.Request().Context(), advertID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot delete advert"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "advert deleted"})
	}
}

func (h *AdvertHandler) UploadImages() echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		files := form.File["images"]
		if len(files) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "no images"})
		}

		urls, err := h.service.UploadImages(files)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot upload images"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"images": urls})
	}
}

func (h *AdvertHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		adverts, err := h.service.GetAll(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "adverts not found"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"adverts": adverts})
	}
}

func (h *AdvertHandler) GetByCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		categoryID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		adverts, err := h.service.GetByCategory(c.Request().Context(), categoryID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "adverts not found"})
		}

		return c.JSON(http.StatusOK, adverts)
	}
}

func (h *AdvertHandler) GetByAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		authorID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		adverts, err := h.service.GetByUserID(c.Request().Context(), authorID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "adverts not found"})
		}

		return c.JSON(http.StatusOK, adverts)
	}
}

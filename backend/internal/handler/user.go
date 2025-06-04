package handler

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(jwtSecret []byte) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.UserRegister
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		token, err := h.userService.Register(c.Request().Context(), jwtSecret, req)
		if err != nil {
			//TODO: handle error
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}
}

func (h *UserHandler) Login(jwtSecret []byte) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.UserLogin
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		token, err := h.userService.Login(c.Request().Context(), jwtSecret, req)
		if err != nil {
			//TODO: handle error
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}
}

func (h *UserHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		user, err := h.userService.GetByID(c.Request().Context(), objID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.UserUpdate
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		id, err := primitive.ObjectIDFromHex(c.Get("userID").(string))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
		}

		err = h.userService.UpdateByID(c.Request().Context(), id, req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "user updated"})
	}
}

package handler

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
	"github.com/labstack/echo/v4"
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

package handler

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/dto"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	fileDir = "./static/images/avatars/"
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

func (h *UserHandler) UploadAvatar() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid, ok := c.Get("userID").(string)
		if !ok || uid == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		}

		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid file"})
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot open file"})
		}
		defer src.Close()

		os.MkdirAll(fileDir, os.ModePerm)
		filename := uid + filepath.Ext(file.Filename)
		dstPath := filepath.Join(fileDir, filename)

		dst, err := os.Create(dstPath)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot create file"})
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot copy file"})
		}

		avatarURL := "/static/images/avatars/" + filename
		if err := h.userService.UpdateAvatar(c.Request().Context(), uid, avatarURL); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot update avatar"})
		}

		return c.JSON(http.StatusOK, map[string]string{"avatarURL": avatarURL})
	}
}

func (h *UserHandler) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid, ok := c.Get("userID").(string)
		if !ok || uid == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		}

		id, err := primitive.ObjectIDFromHex(uid)
		user, err := h.userService.GetByID(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
		}

		return c.JSON(http.StatusOK, user)
	}
}

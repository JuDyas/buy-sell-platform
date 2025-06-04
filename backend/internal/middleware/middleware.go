package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type ctxKey string

const UserIDKey ctxKey = "userID"

func AuthMiddleware(jwtSecret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return next(c)
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header"})
			}

			tokenStr := parts[1]

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token claims"})
			}

			userID, ok := claims["user_id"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid user id"})
			}

			ctx := context.WithValue(c.Request().Context(), UserIDKey, userID)
			c.SetRequest(c.Request().WithContext(ctx))

			c.Set("userID", userID)

			return next(c)
		}
	}

}

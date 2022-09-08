package http

import (
	"github.com/labstack/echo/v4"
	"github.com/quankori/go-clear-architecture/internal/modules/auth"
)

// Map auth routes
func MapAuthRoutes(authGroup *echo.Group, h auth.Handler) {
	authGroup.POST("/register", h.SignUp())
	authGroup.POST("/login", h.SignIn())
}

package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authHttp "github.com/quankori/go-clear-architecture/internal/modules/auth/delivery/http"
	authRepository "github.com/quankori/go-clear-architecture/internal/modules/auth/repository"
	authService "github.com/quankori/go-clear-architecture/internal/modules/auth/services"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	// repos
	userRepo := authRepository.NewUserRepository(s.db, "users")

	//service
	authUC := authService.NewAuthUseCase(userRepo, "Hash_salt", []byte("pokemon"), 86400)

	//handler
	authHandler := authHttp.NewAuthHandler(authUC)

	//middlewares
	// mw := middlewares.NewMiddlewareManager(authUC)

	e.Use(middleware.BodyLimit("2M"))

	//versioning
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandler)

	// If want apply middleware
	// todoHttp.MapAuthRoutes(todoGroup, todoHandler, mw)

	return nil
}

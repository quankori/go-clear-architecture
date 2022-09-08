package middlewares

import "github.com/quankori/go-clear-architecture/internal/modules/auth"

type MiddlewareManager struct {
	authUC auth.UseService
}

func NewMiddlewareManager(authUC auth.UseService) *MiddlewareManager {
	return &MiddlewareManager{authUC}
}

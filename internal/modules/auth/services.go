package auth

import (
	"context"

	"github.com/quankori/go-clear-architecture/internal/models"
)

type UseService interface {
	SignUp(ctx context.Context, username, password string) (*models.User, error)
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, error)
}

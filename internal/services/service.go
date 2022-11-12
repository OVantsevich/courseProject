package services

import (
	"context"
	"github.com/OVantsevich/courseProject/internal/domain"
	"github.com/OVantsevich/courseProject/internal/repository"
)

type Service struct {
	rps repository.UserRepository
}

// NewService create new service connection
func NewService(pool repository.UserRepository) *Service {
	return &Service{rps: pool}
}

func (se *Service) LoginUser(ctx context.Context, login, password string) (string, error) {
	return se.rps.FindUserPasswordByLogin(ctx, login, password)
}

func (se *Service) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	return se.rps.CreateUser(ctx, user)
}

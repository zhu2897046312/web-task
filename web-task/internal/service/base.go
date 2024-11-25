package service

import (
	"web-task/internal/repository"
)

type Service struct {
	repoFactory *repository.RepositoryFactory
}

func NewService(repoFactory *repository.RepositoryFactory) *Service {
	return &Service{
		repoFactory: repoFactory,
	}
} 
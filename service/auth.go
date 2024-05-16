package service

import (
	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/repo"
	"github.com/rs/zerolog"
)

type AuthService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *repo.RepoManager
}

func newAuthService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *AuthService {
	return &AuthService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

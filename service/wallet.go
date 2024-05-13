package service

import (
	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/repo"
	"github.com/rs/zerolog"
)

type WalletService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *repo.RepoManager
}

func newWalletService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *WalletService {
	return &WalletService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

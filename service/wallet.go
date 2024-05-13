package service

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/nano-go/config"
	"github.com/savi2w/nano-go/repo"
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

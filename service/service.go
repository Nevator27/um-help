package service

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/nano-go/config"
	"github.com/savi2w/nano-go/repo"
)

type Service struct {
	User   *UserService
	Wallet *WalletService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{
		User:   newUserService(cfg, logger, repo),
		Wallet: newWalletService(cfg, logger, repo),
	}
}

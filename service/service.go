package service

import (
	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/repo"
	"github.com/rs/zerolog"
)

type Service struct {
	User   *UserService
	Wallet *WalletService
	Auth   *AuthService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{
		User:   newUserService(cfg, logger, repo),
		Wallet: newWalletService(cfg, logger, repo),
		Auth:   newAuthService(cfg, logger, repo),
	}
}

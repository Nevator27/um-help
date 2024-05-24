package service

import (
	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/repo"
	"github.com/Nevator27/um-help/util/cryptoutil"
	"github.com/rs/zerolog"
)

type Service struct {
	User   *UserService
	Wallet *WalletService
	Auth   *AuthService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) (*Service, error) {
	cryptoutil, err := cryptoutil.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		Auth:   newAuthService(cfg, cryptoutil, logger, repo),
		User:   newUserService(cfg, cryptoutil, logger, repo),
		Wallet: newWalletService(cfg, logger, repo),
	}, nil
}

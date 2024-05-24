package service

import (
	"context"
	"strings"

	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/model"
	"github.com/Nevator27/um-help/presenter/req"
	"github.com/Nevator27/um-help/repo"
	"github.com/Nevator27/um-help/util/cryptoutil"
	"github.com/rs/zerolog"
)

type UserService struct {
	config     *config.Config
	cryptoutil *cryptoutil.Cryptoutil
	logger     *zerolog.Logger
	repo       *repo.RepoManager
}

func newUserService(cfg *config.Config, cryptoutil *cryptoutil.Cryptoutil, logger *zerolog.Logger, repo *repo.RepoManager) *UserService {
	return &UserService{
		config:     cfg,
		cryptoutil: cryptoutil,
		logger:     logger,
		repo:       repo,
	}
}

func (s *UserService) New(ctx context.Context, r *req.NewUser) error {
	user := model.User{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		DocumentNumber: r.DocumentNumber,
		Password:       s.cryptoutil.HashPassword(r.Password),
	}

	tx, err := s.repo.MySQL.BeginReadCommittedTx(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	userID, err := s.repo.MySQL.User.Insert(tx, ctx, &user)
	if err != nil {
		return err
	}

	currency, _, err := s.repo.MySQL.Currency.SelectByCurrencyCode(tx, ctx, model.CurrencyBRL)
	if err != nil {
		return err
	}

	w := &model.Wallet{
		OwnerID:    userID,
		CurrencyID: currency.ID,
		Alias:      strings.Join([]string{user.FirstName + "'s", "Wallet"}, " "),
	}

	if err := s.repo.MySQL.Wallet.Insert(tx, ctx, w); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

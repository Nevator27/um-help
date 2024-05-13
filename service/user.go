package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/model"
	"github.com/Nevator27/um-help/presenter/req"
	"github.com/Nevator27/um-help/repo"
	"github.com/rs/zerolog"
)

type UserService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *repo.RepoManager
}

func newUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *UserService {
	return &UserService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

func (s *UserService) New(ctx context.Context, r *req.NewUser) (u *model.User, err error) {
	user := &model.User{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		DocumentNumber: r.DocumentNumber,
	}

	tx, err := s.repo.MySQL.BeginReadCommittedTx(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	userID, err := s.repo.MySQL.User.Insert(tx, ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = userID

	currency, found, err := s.repo.MySQL.Currency.SelectByCurrencyCode(tx, ctx, model.CurrencyBRL)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("cannot find `%s` currency in database", model.CurrencyBRL)
	}

	w := &model.Wallet{
		OwnerID:    user.ID,
		CurrencyID: currency.ID,
		Alias:      strings.Join([]string{user.FirstName + "'s", "Wallet"}, " "),
	}

	if err := s.repo.MySQL.Wallet.Insert(tx, ctx, w); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

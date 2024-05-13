package model

import "time"

type Wallet struct {
	ID         int64     `db:"wallet_id"`
	OwnerID    int64     `db:"owner_id"`
	CurrencyID int64     `db:"currency_id"`
	Alias      string    `db:"alias"`
	Balance    int64     `db:"balance"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

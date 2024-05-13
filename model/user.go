package model

type UserAccount struct {
	ID        int64
	FirstName string
	LastName  string
	CPF       string
	Balance   int64
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type UserTransaction struct {
	ID        int64
	SenderID  int64
	ReciverID int64
	Amount    int64
	CreatedAt string
}

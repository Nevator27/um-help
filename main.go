package main

import (
	"github.com/Nevator27/um-help/routes"
	"github.com/labstack/echo/v4"
)

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

func main() {
	e := echo.New()

	e.POST("/account", routes.HandleCreateAccount)

	e.Logger.Fatal(e.Start(":3000"))
}

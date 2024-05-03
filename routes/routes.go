package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserAccountRequest struct {
	FirstName string `json:"firstName"` //annotation
	LastName  string `json:"lastName"`
	CPF       string `json:"cpf"`
}

func HandleCreateAccount(c echo.Context) error {
	var body UserAccountRequest

	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s", body.FirstName))
}

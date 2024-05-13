package routes

import (
	"fmt"
	"net/http"

	presenter_req "github.com/Nevator27/um-help/presenter/req"
	"github.com/Nevator27/um-help/validations"
	"github.com/labstack/echo/v4"
)

func HandleCreateAccount(c echo.Context) error {
	var body presenter_req.CreateUserAccount

	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	err = validations.ValidateCreateUserRequest(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s", body.FirstName))
}

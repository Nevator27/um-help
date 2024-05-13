package validations

import (
	"errors"

	presenter_req "github.com/Nevator27/um-help/presenter/req"
)

func ValidateCreateUserRequest(req *presenter_req.CreateUserAccount) error {
	if len(req.FirstName) < 3 {
		return errors.New("First name too short")
	}

	if len(req.LastName) < 3 {
		return errors.New("Last name too short")
	}

	if len(req.TaxID) != 11 {
		return errors.New("Tax ID must have 11 characters")
	}

	return nil
}

package validation

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/Nevator27/um-help/presenter/req"
)

func VerifyLoginRequest(rc io.ReadCloser) (r *req.NewUser, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("invalid read closer")
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, errors.New("invalid json payload")
	}

	return r, nil
}

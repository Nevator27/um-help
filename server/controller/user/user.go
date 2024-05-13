package user

import (
	"net/http"

	"github.com/Nevator27/um-help/presenter/res"
	"github.com/Nevator27/um-help/service"
	"github.com/Nevator27/um-help/util/resutil"
	"github.com/Nevator27/um-help/validation"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	resutil *resutil.ResUtil
	svc     *service.Service
}

func New(svc *service.Service, resutil *resutil.ResUtil) *Controller {
	return &Controller{
		resutil: resutil,
		svc:     svc,
	}
}

func (ctrl *Controller) HandleNewUser(ctx echo.Context) error {
	req, err := validation.VerifyNewUserRequest(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	user, err := ctrl.svc.User.New(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	res := &res.User{
		ID: user.ID,
	}

	return ctx.JSON(ctrl.resutil.Wrap(res, nil, http.StatusCreated))
}

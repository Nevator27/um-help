package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/savi2w/nano-go/presenter/res"
	"github.com/savi2w/nano-go/service"
	"github.com/savi2w/nano-go/util/resutil"
	"github.com/savi2w/nano-go/validation"
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

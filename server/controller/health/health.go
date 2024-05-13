package health

import (
	"net/http"

	"github.com/Nevator27/um-help/util/resutil"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	resutil *resutil.ResUtil
}

func New(resutil *resutil.ResUtil) *Controller {
	return &Controller{
		resutil: resutil,
	}
}

func (ctrl *Controller) HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

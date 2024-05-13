package controller

import (
	"github.com/Nevator27/um-help/server/controller/health"
	"github.com/Nevator27/um-help/server/controller/user"
	"github.com/Nevator27/um-help/service"
	"github.com/Nevator27/um-help/util/resutil"
	"github.com/rs/zerolog"
)

type Controller struct {
	HealthController *health.Controller
	UserController   *user.Controller
}

func New(svc *service.Service, logger *zerolog.Logger) *Controller {
	resutil := resutil.New(logger)

	return &Controller{
		HealthController: health.New(resutil),
		UserController:   user.New(svc, resutil),
	}
}

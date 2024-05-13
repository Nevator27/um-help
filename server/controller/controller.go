package controller

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/nano-go/server/controller/health"
	"github.com/savi2w/nano-go/server/controller/user"
	"github.com/savi2w/nano-go/service"
	"github.com/savi2w/nano-go/util/resutil"
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

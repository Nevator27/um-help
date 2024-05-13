package router

import (
	"github.com/labstack/echo/v4"
	"github.com/savi2w/nano-go/config"
	"github.com/savi2w/nano-go/server/controller"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	root.POST("/user", ctrl.UserController.HandleNewUser)
}

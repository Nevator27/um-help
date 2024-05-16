package router

import (
	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/server/controller"
	"github.com/labstack/echo/v4"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	root.POST("/auth/login", ctrl.AuthController.HandleLogin)
	root.POST("/user", ctrl.UserController.HandleNewUser)
}

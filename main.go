package main

import (
	"github.com/Nevator27/um-help/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/account", routes.HandleCreateAccount)

	e.Logger.Fatal(e.Start(":3000"))
}

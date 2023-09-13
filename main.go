package main

import (
	"WebHooks/Routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	Routes.WebhookResponseRoutes(e)
	Routes.DockerHubRoutes(e)
	Routes.QuayHubRoutes(e)
	Routes.HarborRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))

}

package main

import (
	"WebHooks/router"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	router.Routes(e)
	router.WebhookResponseRoutes(e)
	router.QuayHubRoutes(e)
	router.HarborRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))

}

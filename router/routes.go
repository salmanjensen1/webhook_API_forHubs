package router

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {

	//docker-hub routes
	d := e.Group("/dockerhub")
	d.POST("/post-newWebHook", DockerHubCreateWebhook)
	d.DELETE("/delete-webhook", DockerHubDeleteWebhook)
	d.GET("/getTags", DockerHubGetTags)

}

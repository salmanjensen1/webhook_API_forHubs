package router

import (
	"WebHooks/dto"
	"fmt"
	"github.com/labstack/echo/v4"
)

func WebhookResponseRoutes(e *echo.Echo) {
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World. Thy strength befits a crown!")
	//})

	e.POST("/dockerhub-hook", func(c echo.Context) error {
		//binding json data
		var jsonResponse dto.WebhookJSONResponse
		err := c.Bind(&jsonResponse)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return err
		}

		fmt.Println("resp:", jsonResponse)

		fmt.Println("Callback URL:", jsonResponse.CallbackURL)
		fmt.Println("Pushed At:", jsonResponse.PushData.PushedAt)
		fmt.Println("Pusher:", jsonResponse.PushData.Pusher)
		fmt.Println("Tag:", jsonResponse.PushData.Tag)
		fmt.Println("Repository Name:", jsonResponse.Repository.Name)
		fmt.Println("Description:", jsonResponse.Repository.Description)
		return nil
	})

	e.POST("/quay-hook", func(c echo.Context) error {
		var quayResponse dto.QuayImagePushResponse
		err := c.Bind(&quayResponse)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return err
		}

		fmt.Println("Name:", quayResponse.Name)
		fmt.Println("Repository:", quayResponse.Repository)
		fmt.Println("Namespace:", quayResponse.Namespace)
		fmt.Println("Docker URL:", quayResponse.DockerURL)
		fmt.Println("Homepage:", quayResponse.Homepage)
		fmt.Println("Updated Tags:", quayResponse.UpdatedTags)
		return nil

	})

	e.POST("/harbor-hook", func(c echo.Context) error {
		var harborResponse interface{}
		err := c.Bind(&harborResponse)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return err
		}

		fmt.Println("Harbor Response:", harborResponse)
		//fmt.Println("Name:", harborResponse.Name)
		//fmt.Println("Repository:", harborResponse.Repository)
		//fmt.Println("Namespace:", harborResponse.Namespace)
		//fmt.Println("Docker URL:", harborResponse.DockerURL)
		//fmt.Println("Homepage:", harborResponse.Homepage)
		//fmt.Println("Updated Tags:", harborResponse.UpdatedTags)
		return nil

	})
}

package Routes

import (
	"WebHooks/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Webhook struct {
	Name    string `json:"name"`
	HookURL string `json:"hook_url"`
}

type RequestData struct {
	Name                string    `json:"name"`
	ExpectFinalCallback bool      `json:"expect_final_callback"`
	Webhooks            []Webhook `json:"webhooks"`
}

func DockerHubRoutes(e *echo.Echo) {
	e.POST("/post-newWebHook", func(c echo.Context) error {
		var token = c.Request().Header.Get("Authorization")
		url := "https://hub.docker.com/v2/repositories/salmanjensen/klovercloud_dashboard/webhook_pipeline/"
		var inputData RequestData
		err := c.Bind(&inputData)
		if err != nil {
			return err
		}
		fmt.Println(inputData)

		// Create a JSON object
		//inputData := RequestData{
		//	Name:                "test5",
		//	ExpectFinalCallback: false,
		//	Webhooks: []Webhook{
		//		{
		//			Name:    "test5",
		//			HookURL: "https://969e-103-217-110-241.ngrok-free.app/test4",
		//		},
		//	},
		//}

		// Convert JSON object to byte data
		data, err := json.Marshal(inputData)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return err
		}

		fmt.Println("Data:", data)

		// Create a new request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}

		// Set the Content-Type header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		// Perform the HTTP request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return err
		}
		defer resp.Body.Close()

		// Read and print the response
		body := new(bytes.Buffer)
		_, _ = body.ReadFrom(resp.Body)
		fmt.Println("Response:", resp.Status)
		fmt.Println("Response Body:", body.String())
		return nil
	})
	e.DELETE("/delete-webhook/:id", func(c echo.Context) error {
		url := "https://hub.docker.com/v2/repositories/salmanjensen/klovercloud_dashboard/webhook_pipeline/"
		id := c.Param("id")

		req, err := http.NewRequest("DELETE", url+id, nil)
		if err != nil {
			fmt.Println("Error creating delete request:", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+config.AuthToken)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making delete webhook:", err)
			return err
		}

		defer resp.Body.Close()

		// Read and print the response
		body := new(bytes.Buffer)
		_, _ = body.ReadFrom(resp.Body)
		fmt.Println("Response:", resp.Status)
		fmt.Println("Response Body:", body.String())
		return nil
	})
}

package Routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type HarborRequestDto struct {
	Enabled    bool     `json:"enabled"`
	EventTypes []string `json:"event_types"`
	Targets    []struct {
		Type           string `json:"type"`
		Address        string `json:"address"`
		SkipCertVerify bool   `json:"skip_cert_verify"`
		PayloadFormat  string `json:"payload_format"`
		AuthHeader     string `json:"auth_header"`
	} `json:"targets"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func HarborRoutes(e *echo.Echo) {
	h := e.Group("/harbor")
	h.POST("/post-newWebHook", func(c echo.Context) error {
		var authHeader = c.Request().Header.Get("Authorization")
		url := "https://core-harbor.kc-cp.klovercloud.io/api/v2.0/projects/2/webhook/policies"
		var inputData HarborRequestDto
		err := c.Bind(&inputData)
		if err != nil {
			return err
		}
		fmt.Println(inputData)

		//extract info from header

		if authHeader == "" {
			fmt.Println("Authorization header is missing")
		}

		// Check if the Authorization header starts with "Basic "
		if !strings.HasPrefix(authHeader, "Basic ") {
			fmt.Println("Invalid Authorization method")
		}

		// Extract and decode the Base64-encoded part of the header
		authValue := strings.TrimPrefix(authHeader, "Basic ")
		fmt.Println("This is the auth value:", authValue)

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
		req.Header.Set("Authorization", authHeader)

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
}

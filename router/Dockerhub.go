package router

import (
	"WebHooks/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DockerHubGetAccessToken(c echo.Context) error {
	url := "https://hub.docker.com/v2/users/login/"
	var inputData dto.DockerHubLoginRequest
	err := c.Bind(&inputData)
	if err != nil {
		return err
	}

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
}

func DockerHubCreateWebhook(c echo.Context) error {
	var token = c.Request().Header.Get("Authorization")
	url := "https://hub.docker.com/v2/repositories/salmanjensen/klovercloud_dashboard/webhook_pipeline/"
	var inputData dto.WebhookRequestData
	err := c.Bind(&inputData)
	if err != nil {
		return err
	}
	fmt.Println(inputData)
	fmt.Println("This is the token:", token)

	// Convert JSON object to byte data
	data, err := json.Marshal(inputData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

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
}

func DockerHubDeleteWebhook(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	url := "https://hub.docker.com/v2/repositories/salmanjensen/klovercloud_dashboard/webhook_pipeline/"
	id := c.Param("id")

	req, err := http.NewRequest("DELETE", url+id, nil)
	if err != nil {
		fmt.Println("Error creating delete request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)
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
}

func DockerHubGetTags(c echo.Context) error {
	var token = c.Request().Header.Get("Authorization")
	url := "https://hub.docker.com/v2/repositories/salmanjensen/klovercloud_dashboard/tags"

	req, err := http.NewRequest("GET", url, nil)
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
	var response dto.TagResponse

	body := new(bytes.Buffer)

	_, _ = body.ReadFrom(resp.Body)
	fmt.Println("Response:", resp.Status)
	fmt.Println("Response Body:", body.String())

	// Unmarshal the JSON response into the DTO struct
	if err := json.Unmarshal(body.Bytes(), &response); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return err
	}

	// Now you can access the data in the response struct
	fmt.Println("Count:", response.Count)
	fmt.Println("Results:")
	for _, result := range response.Results {
		fmt.Println("Name:", result.Name)
		fmt.Println("Last Updated:", result.LastUpdated)
		fmt.Println("Last Updater Username:", result.LastUpdaterUsername)
	}

	return nil
}

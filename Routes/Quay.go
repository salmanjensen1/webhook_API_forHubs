package Routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type QuayRequestDto struct {
	Event  string `json:"event"`
	Method string `json:"method"`
	Config struct {
		URL string `json:"url"`
	} `json:"config"`
	EventConfig interface{} `json:"eventConfig"`
}

func QuayHubRoutes(e *echo.Echo) {

	q := e.Group("/quay")
	q.POST("/post-newWebHook", func(c echo.Context) error {
		cookieStr := c.Request().Header.Get("Cookie")
		x_csrf_Token := c.Request().Header.Get("X-Csrf-Token")
		fmt.Println("Cookie: ", cookieStr)
		fmt.Println(c.Cookie("_csrf_token"))
		url := "https://quay.io/api/v0/repository/hasanshoron/webhook-test/notification/"
		var inputData QuayRequestDto
		err := c.Bind(&inputData)
		if err != nil {
			return err
		}
		fmt.Println(inputData)

		// Convert JSON object to byte data
		data, err := json.Marshal(inputData)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return err
		}
		cookieName := "_csrf_token" // parsing the cookie with this name to find the value
		cookieValue := extractValueFromCookie(cookieStr, cookieName)
		if cookieValue == "" {
			fmt.Println("Could not find cookie value in the string.")
		}

		// Create a new request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}

		cookie := http.Cookie{}

		cookie.Name = cookieName
		cookie.Value = cookieValue
		cookie.Path = "/"
		cookie.Domain = "quay.io"

		// Set the Content-Type header
		req.Header.Set("Content-Type", "application/json")
		req.AddCookie(&cookie)
		req.Header.Add("X-Csrf-Token", x_csrf_Token)

		//print request header
		fmt.Println("Request Header:")
		for name, value := range req.Header {
			fmt.Printf("%s: %s\n", name, value)
		}

		// Perform the HTTP request
		client := &http.Client{}
		fmt.Println("Request:", req)
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
	q.DELETE("/delete-webhook/:id", func(c echo.Context) error {
		cookieStr := c.Request().Header.Get("Cookie")
		x_csrf_Token := c.Request().Header.Get("X-Csrf-Token")
		fmt.Println("Cookie: ", cookieStr)
		fmt.Println(c.Cookie("_csrf_token"))
		url := "https://quay.io/api/v0/repository/hasanshoron/webhook-test/notification/"
		id := c.Param("id")
		cookieName := "_csrf_token" // Modify this to the desired cookie name
		cookieValue := extractValueFromCookie(cookieStr, cookieName)
		if cookieValue == "" {
			fmt.Println("Could not find cookie value in the string.")
		}

		req, err := http.NewRequest("DELETE", url+id, nil)
		if err != nil {
			fmt.Println("Error creating delete request:", err)
			return err
		}
		cookie := http.Cookie{}

		cookie.Name = cookieName
		cookie.Value = cookieValue
		cookie.Path = "/"
		cookie.Domain = "quay.io"

		// Set the Content-Type header
		req.Header.Set("Content-Type", "application/json")
		req.AddCookie(&cookie)
		req.Header.Add("X-Csrf-Token", x_csrf_Token)
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

func extractValueFromCookie(cookieStr string, cookieName string) string {
	parts := strings.Split(cookieStr, "; ")
	for _, part := range parts {
		if strings.HasPrefix(part, cookieName+"=") {
			return strings.TrimPrefix(part, cookieName+"=")
		}
	}
	return ""
}

package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type WebhookCreateDTO struct {
	UUID   string `json:"uuid"`
	Event  string `json:"event"`
	Method string `json:"method"`
	Config struct {
		URL string `json:"url"`
	} `json:"config"`
	EventConfig interface{} `json:"eventConfig"`
}

type RepositoryInfo struct {
	Username string `json:"username"`
	RepoName string `json:"repo_name"`
}

type TagInfo struct {
	Name           string `json:"name"`
	Reversion      bool   `json:"reversion"`
	StartTS        int64  `json:"start_ts"`
	ManifestDigest string `json:"manifest_digest"`
	IsManifestList bool   `json:"is_manifest_list"`
	Size           int    `json:"size"`
	LastModified   string `json:"last_modified"`
}

type ImageListResponseFromQuay struct {
	Tags          []TagInfo `json:"tags"`
	Page          int       `json:"page"`
	HasAdditional bool      `json:"has_additional"`
}

type ImageListResponse struct {
	ImageCount int      `json:"image_count"`
	ImageNames []string `json:"image_names"`
}

type Repository struct {
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
	Kind        string `json:"kind"`
	State       string `json:"state"`
	IsStarred   bool   `json:"is_starred"`
}

type RepositoryList struct {
	Repositories []Repository `json:"repositories"`
}

type RepositoryListResponse struct {
	RepositoryCount int      `json:"repository_count"`
	RepositoryName  []string `json:"repositories"`
}

func QuayHubRoutes(e *echo.Echo) {

	q := e.Group("/quay")

	q.GET("/get-repositories", func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		url := "https://quay.io/api/v1/repository?"
		public := c.QueryParam("public")
		private := c.QueryParam("private")
		namespace := c.QueryParam("namespace")

		if public != "" {
			url = url + "public=" + public + "&"
		}
		if private != "" {
			url = url + "private=" + private + "&"
		}
		if namespace != "" {
			url = url + "namespace=" + namespace
		}

		fmt.Println("URL:", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return err
		}
		defer resp.Body.Close()

		// Read and print the response
		var data RepositoryList
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			fmt.Println("Error parsing JSON:", err.Error())
		}

		repositoryCount := len(data.Repositories)
		var repositoryNames []string

		for _, repo := range data.Repositories {
			repositoryNames = append(repositoryNames, repo.Name)
			fmt.Println(repo.Name)
		}

		repositoryListResponse := RepositoryListResponse{
			RepositoryCount: repositoryCount,
			RepositoryName:  repositoryNames,
		}

		return c.JSON(http.StatusOK, repositoryListResponse)
	})

	q.POST("/get-tags", func(c echo.Context) error {
		var repositoryInfo RepositoryInfo
		err := c.Bind(&repositoryInfo)
		fmt.Println("repositoryInfo:", repositoryInfo)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err.Error())
		}

		token := c.Request().Header.Get("Authorization")
		url := "https://quay.io/api/v1/repository/" + repositoryInfo.Username + "/" + repositoryInfo.RepoName + "/tag/" + "?"
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")
		onlyActiveTags := c.QueryParam("onlyActiveTags")

		if limit != "" {
			url = url + "limit=" + limit
		}
		if page != "" {
			url = url + "page=" + page
		}
		if onlyActiveTags != "" {
			url = url + "onlyActiveTags=" + onlyActiveTags
		}
		fmt.Println("url:", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return err
		}
		defer resp.Body.Close()

		// Create an instance of your response struct
		var data ImageListResponseFromQuay

		// Decode the response body into the struct
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&data); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		imageCount := len(data.Tags)
		var tagNames []string
		// Read and print the response
		for _, tag := range data.Tags {
			tagNames = append(tagNames, tag.Name)
			fmt.Println("Repository Name:", tag.Name)
		}
		imageListResponse := ImageListResponse{
			ImageNames: tagNames,
			ImageCount: imageCount,
		}
		return c.JSON(http.StatusOK, imageListResponse)
	})

	q.POST("/webhook", func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		url := "https://quay.io/api/v1/repository/hasanshoron/webhook-test/notification/"
		var inputData WebhookCreateDTO
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

		// Create a new request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}

		// Set the Content-Type header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)

		// Perform the HTTP request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return err
		}
		defer resp.Body.Close()

		var webhookCreate WebhookCreateDTO
		err = json.NewDecoder(resp.Body).Decode(&webhookCreate)
		if err != nil {
			fmt.Println("Error parsing JSON:", err.Error())
		}
		return c.JSON(http.StatusCreated, webhookCreate)
	})
	q.DELETE("/webhook/:id", func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		url := "https://quay.io/api/v1/repository/hasanshoron/webhook-test/notification/"
		id := c.Param("id")

		req, err := http.NewRequest("DELETE", url+id, nil)
		if err != nil {
			fmt.Println("Error creating delete request:", err)
			return err
		}

		// Set the Content-Type header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making delete webhook:", err)
			return err
		}

		defer resp.Body.Close()

		// Read and print the response
		fmt.Println("Response:", resp.Status)

		if resp.Status == "204 No Content" {
			return c.JSON(http.StatusOK, "Webhook deleted successfully")
		} else {
			return c.JSON(http.StatusBadRequest, "Webhook not deleted")
		}
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

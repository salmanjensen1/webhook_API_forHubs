package dto

import "time"

type PushData struct {
	PushedAt int64  `json:"pushed_at"`
	Pusher   string `json:"pusher"`
	Tag      string `json:"tag"`
}

type Repository struct {
	CommentCount    int    `json:"comment_count"`
	DateCreated     int64  `json:"date_created"`
	Description     string `json:"description"`
	Dockerfile      string `json:"dockerfile"`
	FullDescription string `json:"full_description"`
	IsOfficial      bool   `json:"is_official"`
	IsPrivate       bool   `json:"is_private"`
	IsTrusted       bool   `json:"is_trusted"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	Owner           string `json:"owner"`
	RepoName        string `json:"repo_name"`
	RepoURL         string `json:"repo_url"`
	StarCount       int    `json:"star_count"`
	Status          string `json:"status"`
}

type WebhookJSONResponse struct {
	CallbackURL string     `json:"callback_url"`
	PushData    PushData   `json:"push_data"`
	Repository  Repository `json:"repository"`
}

type TagResponse struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []TagResult `json:"results"`
}

type TagResult struct {
	Creator             int         `json:"creator"`
	ID                  int         `json:"id"`
	Images              []TagImage  `json:"images"`
	LastUpdated         time.Time   `json:"last_updated"`
	LastUpdater         int         `json:"last_updater"`
	LastUpdaterUsername string      `json:"last_updater_username"`
	Name                string      `json:"name"`
	Repository          int         `json:"repository"`
	FullSize            int         `json:"full_size"`
	V2                  bool        `json:"v2"`
	TagStatus           string      `json:"tag_status"`
	TagLastPulled       interface{} `json:"tag_last_pulled"`
	TagLastPushed       time.Time   `json:"tag_last_pushed"`
	MediaType           string      `json:"media_type"`
	ContentType         string      `json:"content_type"`
	Digest              string      `json:"digest"`
}

type TagImage struct {
	Architecture string      `json:"architecture"`
	Features     string      `json:"features"`
	Variant      interface{} `json:"variant"`
	Digest       string      `json:"digest"`
	OS           string      `json:"os"`
	OSFeatures   string      `json:"os_features"`
	OSVersion    interface{} `json:"os_version"`
	Size         int         `json:"size"`
	Status       string      `json:"status"`
	LastPulled   interface{} `json:"last_pulled"`
	LastPushed   time.Time   `json:"last_pushed"`
}

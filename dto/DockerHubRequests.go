package dto

type Webhook struct {
	Name    string `json:"name"`
	HookURL string `json:"hook_url"`
}

type WebhookRequestData struct {
	Name                string    `json:"name"`
	ExpectFinalCallback bool      `json:"expect_final_callback"`
	Webhooks            []Webhook `json:"webhooks"`
}

type DockerHubLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

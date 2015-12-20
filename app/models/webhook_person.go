package models

type Webhook struct {
	Result    Person  `json:"result"`
	WebhookId string   `json:"webhookId"`
}
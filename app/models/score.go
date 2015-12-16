package models

type Score struct {
	Type     string `json:"type"`
	Value    int    `json:"value"`
	Provider string `json:"provider"`
}

package models

type County struct {
	Location
	Code string `json:"code"`
}

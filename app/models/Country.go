package models


type Country struct {
	Location
	Code string `json:"code"`
}
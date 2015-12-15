package models


type DigitalFootprint struct {
	Topics []Topic `json:"topics"`
	Scores []Score `json:"scores"`
}
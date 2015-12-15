package models


type Demographics struct {
	LocationGeneral string          `json:"locationGeneral"`
	LocationDeduced LocationDeduced `json:"locationDeduced"`
	Age             LocationDeduced `json:"age"`
	Gender          LocationDeduced `json:"gender"`
	AgeRange        LocationDeduced `json:"ageRange"`
}
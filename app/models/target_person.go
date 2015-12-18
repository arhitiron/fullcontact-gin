package models

type TargetPerson struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Processed bool   `json:"processed"`
}

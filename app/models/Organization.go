package models


type Organization struct {
	Title     string    `json:"title"`
	Name      string    `json:"name"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	IsPrimary bool      `json:"isPrimary"`
	Current   bool      `json:"current"`

}

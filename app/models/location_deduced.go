package models

type LocationDeduced struct {
	NormalizedLocation string    `json:"normalizedLocation"`
	DeducedLocation    string    `json:"deducedLocation"`
	City               City      `json:"city"`
	State              State     `json:"state"`
	Country            Country   `json:"country"`
	Continent          Continent `json:"continent"`
	County             County    `json:"county"`
	Likelihood         string    `json:"likelihood"`
}

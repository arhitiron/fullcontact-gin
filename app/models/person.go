package models

type Person struct {
	Id               int
	Status           int              `json:"status"`
	RequestId        string           `json:"requestId"`
	Likelihood       int              `json:"likelihood"`
	ContactInfo      ContactInfo      `json:"contactInfo"`
	Demographics     Demographics     `json:"demographics"`
	Photos           []Photo          `json:"photos"`
	SocialProfiles   []SocialProfile  `json:"socialProfiles"`
	DigitalFootprint DigitalFootprint `json:"digitalFootprint"`
	Organizations    []Organization   `json:"organizations"`
}

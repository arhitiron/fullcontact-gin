package models


type ContactInfo struct {
	FamilyName  string    `json:"familyName"`
	GivenName   string    `json:"givenName"`
	FullName    string    `json:"fullName"`
	MiddleNames []string  `json:"middleNames"`
	Websites    []Website `json:"websites"`
	Chats       []Chat    `json:"chats"`
}
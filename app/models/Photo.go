package models

type Photo struct {
	TypeId    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	Url       string `json:"url"`
	IsPrimary bool   `json:"isPrimary"`
}

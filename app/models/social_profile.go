package models

type SocialProfile struct {
	TypeId    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	Id        string `json:"id"`
	Username  string `json:"username"`
	Url       string `json:"url"`
	Bio       string `json:"bio"`
	Rss       string `json:"rss"`
	Following int    `json:"following"`
	Followers int    `json:"followers"`
}

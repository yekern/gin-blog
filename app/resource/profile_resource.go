package resource

type Profile struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type Role struct {
	Name string `json:"name"`
}

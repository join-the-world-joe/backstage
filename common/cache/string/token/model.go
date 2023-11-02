package token

type Model struct {
	UserId int64  `json:"user_id"`
	Secret string `json:"secret"`
}

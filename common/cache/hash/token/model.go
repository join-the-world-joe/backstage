package token

type Token struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	UserId      string `json:"user_id"`
}

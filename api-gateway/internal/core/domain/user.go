package domain

type User struct {
	Uuid      string `json:"uuid"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AuthToken string `json:"auth_token"`
}

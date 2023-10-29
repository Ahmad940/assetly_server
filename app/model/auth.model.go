package model

type Login struct {
	CountryCode string `json:"country_code" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	PassCode    string `json:"pass_code" validate:"len=6,required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

package models

type AuthUserRequest struct {
	PersonalNumber string `json:"personal_number" binding:"required,len=7"`
	SecretCode     string `json:"secret_code" binding:"required,len=4"`
}

package models

type CreateUserRequest struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	PersonalNumber string `json:"personal_number" binding:"required,len=-7"`
	SecretCode     string `json:"secret_code" binding:"required,len=4"`
	PhoneNumber    string `json:"phone_number" binding:"required,len=10"`
}

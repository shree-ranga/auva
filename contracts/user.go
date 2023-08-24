package contracts

import (
	"mime/multipart"
	"time"
)

type RegisterUserRequest struct {
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	DOB       time.Time      `json:"dob"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone_number"`
	CV        multipart.File `json:"cv"`
}
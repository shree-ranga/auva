package contracts

import (
	"mime/multipart"
	"net/url"
	"regexp"
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

func (r *RegisterUserRequest) Validate() url.Values {
	err := url.Values{}

	phoneNumberRegex := regexp.MustCompile(`^[0-9]+$`)
	if !phoneNumberRegex.MatchString(r.Phone) {
		err.Add("phone_number", "Phone number should only contain numbers")
	}

	// Validate phone number length
	// assuming india is the only country
	if len(r.Phone) > 10 {
		err.Add("phone_number", "Phone number should not exceed 10 digits")
	}
	return err
}

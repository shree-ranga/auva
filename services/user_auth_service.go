package services

import (
	"auva/contracts"
	"auva/domains"
	"auva/repo"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-gomail/gomail"
)

type UserAuthService interface {
	RegisterUser(ctx context.Context, body contracts.RegisterUserRequest) error
}

type userAuthService struct {
	userAuthRepo repo.UserAuthRepo
}

func NewUserAuthService(repo repo.UserAuthRepo) UserAuthService {
	return &userAuthService{
		userAuthRepo: repo,
	}
}

func (s *userAuthService) RegisterUser(ctx context.Context, body contracts.RegisterUserRequest) error {
	var err error

	// saving resume files to the local disk.
	// ideally should be stored in AWS S3.
	// converting .doc and .docx file to pdf as well for simplicity.
	cvFilePath := fmt.Sprintf("/Users/srr/Desktop/auva/media/%s_resume.pdf", body.LastName)
	cvFileOut, err := os.Create(cvFilePath)
	if err != nil {
		return err
	}
	defer cvFileOut.Close()

	// Copy the contents of the uploaded CV to the local file
	_, err = io.Copy(cvFileOut, body.CV)
	if err != nil {
		return err
	}

	domain := domains.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		DOB:       body.DOB,
		Email:     body.Email,
		Phone: domains.PhoneNumber{
			Number: body.Phone,
		},
		CVPath: cvFilePath,
	}

	err = s.userAuthRepo.Insert(ctx, domain)

	if err == nil {
		// send email asynchrnously
		go sendEmailAsync(body.Email)
	}

	return err
}

func sendEmailAsync(emailAddress string) {
	m := gomail.NewMessage()

	smtpUrl := "smtp.example.com"
	smtpPort := 587
	companyEmail := "no-reply@auva.com"
	companyPassword := "your-password"
	subject := "Registration Complete"
	content := "Thank you for registering!"
	m.SetHeader("From", companyEmail)
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(smtpUrl, smtpPort, companyEmail, companyPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Println("could not send email")
	}
}

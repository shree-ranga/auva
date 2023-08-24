package services

import (
	"auva/contracts"
	"auva/domains"
	"auva/repo"
	"context"
	"fmt"
	"io"
	"os"
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

	return s.userAuthRepo.Insert(ctx, domain)
}

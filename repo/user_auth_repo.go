package repo

import (
	"auva/domains"
	"context"
	"log"
	"time"

	"github.com/edgedb/edgedb-go"
)

type UserAuthRepo interface {
	Insert(ctx context.Context, model domains.User) error
}

type userAuthRepo struct {
	db *edgedb.Client
}

func NewAuthRepo(db *edgedb.Client) UserAuthRepo {
	return &userAuthRepo{
		db: db,
	}
}

func (s *userAuthRepo) Insert(ctx context.Context, model domains.User) error {
	dob := time.Date(model.DOB.Year(), model.DOB.Month(), model.DOB.Day(), 0, 0, 0, 0, time.UTC)
	var inserted struct{ id edgedb.UUID }
	err := s.db.QuerySingle(ctx, `INSERT User {
		first_name := <str>$0,
		last_name := <str>$1,
		date_of_birth := <datetime>$2,
		email := <str>$3,
		phone_number := (INSERT PhoneNumber {
			number := <str>$4
		}),
		cv_file_path := <str>$5
	};`, &inserted, model.FirstName, model.LastName, dob, model.Email, model.Phone.Number, model.CVPath)

	if err != nil {
		log.Fatal("error inserting the data", err.Error())
	}
	return err
}

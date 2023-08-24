package domains

import (
	"time"

	"github.com/edgedb/edgedb-go"
)

type User struct {
	ID        edgedb.UUID `edgedb:"id"`
	FirstName string      `edgedb:"first_name"`
	LastName  string      `edgedb:"last_name"`
	DOB       time.Time   `edgedb:"date_of_birth"`
	Email     string      `edgedb:"email"`
	Phone     PhoneNumber `edgedb:"phone_number"`
	CVPath    string      `edgedb:"cv_file_path"`
}

type PhoneNumber struct {
	Number string `edgedb:"number"`
}

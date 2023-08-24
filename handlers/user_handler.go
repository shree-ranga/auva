package handlers

import (
	"auva/contracts"
	"auva/services"
	"context"
	"net/http"
	"time"
)

func RegisterUser(service services.UserAuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		var req contracts.RegisterUserRequest
		req.FirstName = r.FormValue("first_name")
		req.LastName = r.FormValue("last_name")
		req.Email = r.FormValue("email")
		req.Phone = r.FormValue("phone_number")
		cvFile, _, err := r.FormFile("cv")
		if err != nil {
			http.Error(w, "Error retrieving the CV file", http.StatusBadRequest)
			return
		}
		defer cvFile.Close()
		req.CV = cvFile
		dobStr := r.FormValue("dob")
		dob, err := time.Parse("2006-01-02", dobStr)
		if err != nil {
			http.Error(w, "Invalid date of birth format", http.StatusBadRequest)
			return
		}
		req.DOB = dob

		// add contract validation here

		// pass the request to service layer
		err = service.RegisterUser(ctx, req)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User registered successfully!"))
	}
}

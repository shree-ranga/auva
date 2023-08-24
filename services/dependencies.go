package services

import (
	"auva/appContext"
	"auva/repo"
)

type ServerDependencies struct {
	AuthService UserAuthService
}

func InstantiateServerDependencies() *ServerDependencies {
	db := appContext.GetDB()
	authRepo := repo.NewAuthRepo(db)
	authService := NewUserAuthService(authRepo)

	return &ServerDependencies{
		AuthService: authService,
	}
}

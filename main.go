package main

import (
	"auva/appContext"
	"auva/router"
	"auva/services"
	"fmt"
	"net/http"
)

func main() {
	// app context intialization
	appContext.Init()

	// router intialization
	dependencies := services.InstantiateServerDependencies()
	r := router.InitServerRouter(*dependencies)

	// new http server
	httpServer := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	fmt.Println("Server is starting on port 8000...")
	if err := httpServer.ListenAndServe(); err != nil {
		panic("http server startup failed")
	}
}

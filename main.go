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
	// ctx := context.Background()

	// opts := edgedb.Options{
	// 	Concurrency: 4,
	// }
	// client, err := edgedb.CreateClient(ctx, opts)
	// if err != nil {
	// 	panic("could not connect to db")
	// }
	// // var inserted struct{ id edgedb.UUID }
	// err = client.Execute(ctx, `INSERT User {
	// 	first_name := <str>$0,
	// 	last_name := <str>$1,
	// 	date_of_birth := <datetime>$2,
	// 	email := <str>$3,
	// 	phone_number := (INSERT PhoneNumber {
	// 		number := <str>$4
	// 	}),
	// 	cv_file_path := <str>$5
	// };`, "Shree", "Ranga", time.Date(1984, 3, 1, 0, 0, 0, 0, time.UTC), "raju@sr.com", "9898989", "shre/sr.pdf")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
}

package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"todo_app/api"
)

func main() {
	//Create API and Initialize Routes, DB Store,Validator & JWT
	r := api.NewAPI()
	r.Initialize()

	//Middleware
	n := negroni.New()
	//n.UseHandler(middleware.NewAuthorization(r.MainRouter))
	n.UseHandler(r.Authentication)

	log.Fatal(http.ListenAndServe(":8080", n))
}

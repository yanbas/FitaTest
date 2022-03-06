package main

import (
	"fita/DB"
	"fita/graphql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init DB
	err := DB.DbInit()
	if err != nil {
		log.Fatal(err)
	}

	h, err := graphql.NewHandler()
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	r.Methods("POST").Path("/").Handler(h)
	http.Handle("/", r)
	http.ListenAndServe(":5100", nil)
}
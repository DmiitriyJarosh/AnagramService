package main

import (
	"./controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/load",
		controllers.LoadWords).Methods("POST")
	router.HandleFunc("/get",
		controllers.GetAnagramsFor).Methods("GET")

	port := "8080"
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal(err)
	}
}

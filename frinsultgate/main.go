package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	if envPort, present := os.LookupEnv("PORT"); present {
		port = envPort
	}
	r := mux.NewRouter()

	r.Methods("GET").Path("/insults/{id:[0-9]+}").HandlerFunc(getInsultById)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

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
	r.Methods("POST").Path("/insults/vote/{id:[0-9]+}").HandlerFunc(voteInsultById)
	r.Methods("DELETE").Path("/insults/{id:[0-9]+}").HandlerFunc(deleteInsultById)
	r.Methods("PATCH").Path("/insults/{id:[0-9]+}").HandlerFunc(updateInsultById)

	r.Methods("PUT").Path("/insults").HandlerFunc(putInsult)

	r.Methods("GET").Path("/insults").HandlerFunc(getInsults)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pijalu/go.hands.two/env"
)

func main() {
	port := env.GetEnvWithDefault("PORT", "8080")
	r := mux.NewRouter()

	r.Methods("GET").Path("/insults/{id:[0-9]+}").HandlerFunc(getInsultByID)
	r.Methods("POST").Path("/insults/vote/{id:[0-9]+}").HandlerFunc(voteInsultByID)
	r.Methods("DELETE").Path("/insults/{id:[0-9]+}").HandlerFunc(deleteInsultByID)
	r.Methods("PATCH").Path("/insults/{id:[0-9]+}").HandlerFunc(updateInsultByID)

	r.Methods("PUT").Path("/insults").HandlerFunc(putInsult)

	r.Methods("GET").Path("/insults").HandlerFunc(getInsults)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

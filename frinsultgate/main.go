package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	micro "github.com/pijalu/go-micro"
	"github.com/pijalu/go.hands.two/env"
	"github.com/pijalu/go.hands.two/frinsultproto"

	// Register k8s specific
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/selector/static"
)

var friService frinsultproto.FrinsultService

func main() {
	// Load micro client
	service := micro.NewService(
		micro.Name("frinsult.client.micro"),
	)
	service.Init()
	friService = frinsultproto.NewFrinsultService("frinsult.srv.micro", service.Client())

	// start gateway
	port := env.GetEnvWithDefault("PORT", "8080")
	r := mux.NewRouter()

	r.Methods("GET").Path("/insults/{id:[0-9]+}").HandlerFunc(getInsultByID)
	r.Methods("DELETE").Path("/insults/{id:[0-9]+}").HandlerFunc(deleteInsultByID)
	r.Methods("PATCH").Path("/insults/{id:[0-9]+}").HandlerFunc(updateInsultByID)
	r.Methods("PUT").Path("/insults").HandlerFunc(putInsult)

	r.Methods("POST").Path("/insults/upvote/{id:[0-9]+}").HandlerFunc(upvoteInsultByID)
	r.Methods("POST").Path("/insults/downvote/{id:[0-9]+}").HandlerFunc(downvoteInsultByID)

	r.Methods("GET").Path("/insults").HandlerFunc(getInsults)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

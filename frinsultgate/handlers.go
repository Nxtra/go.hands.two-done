package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func replyWithJson(w http.ResponseWriter, reply interface{}, status int) {
	b, err := json.Marshal(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Error during marshal: %e", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

func getIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		return 0, errors.New("No id provided")
	}

	return strconv.Atoi(id)
}

func getInsults(w http.ResponseWriter, r *http.Request) {
	replyWithJson(w, &[]insult{{
		ID:    1,
		Text:  "Lazy fuck !",
		Score: 69,
	}}, http.StatusOK)
}

func deleteInsultById(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve ID: " + err.Error()))
		return
	}

	// TODO

	log.Printf("Deleted insults %d", id)
	w.WriteHeader(http.StatusOK)
}

func getInsultById(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve ID: " + err.Error()))
		return
	}

	// TODO

	replyWithJson(w, &insult{
		ID:   id,
		Text: "Fuck off !",
	}, http.StatusOK)
}

func putInsult(w http.ResponseWriter, r *http.Request) {
	log.Printf("Added insults !")
	decoder := json.NewDecoder(r.Body)

	entity := insult{}
	if err := decoder.Decode(&entity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request:" + err.Error()))
		return
	}
	// TODO
	w.WriteHeader(http.StatusOK)
}

func voteInsultById(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO

	log.Printf("Voted for insult %d", id)
	w.WriteHeader(http.StatusOK)
}

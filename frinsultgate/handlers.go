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

func replyWithJSON(w http.ResponseWriter, reply interface{}, status int) {
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
	replyWithJSON(w, &[]insult{{
		ID:    1,
		Text:  "Lazy fuck !",
		Score: 69,
	}}, http.StatusOK)
}

func deleteInsultByID(w http.ResponseWriter, r *http.Request) {
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

func getInsultByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve ID: " + err.Error()))
		return
	}

	// TODO

	replyWithJSON(w, &insult{
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

	log.Printf("Putting %s", entity.String())
	// TODO
	w.WriteHeader(http.StatusOK)
}

func updateInsultByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve ID: " + err.Error()))
		return
	}

	decoder := json.NewDecoder(r.Body)
	entity := insult{}
	if err := decoder.Decode(&entity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request:" + err.Error()))
		return
	}
	entity.ID = id

	log.Printf("Updating entity %T", entity)
	// TODO
	w.WriteHeader(http.StatusOK)
}

func voteInsultByID(w http.ResponseWriter, r *http.Request) {
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

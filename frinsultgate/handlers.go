package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

func getInsultById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("gimme an id you moron !"))
		return
	}

	// TODO

	replyWithJson(w, &insult{
		ID:   69,
		Text: "Fuck off !",
	}, http.StatusOK)
}

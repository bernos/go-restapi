package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, todos, http.StatusOK, handleError(w))
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		sendJSON(w, err, 422, handleError(w))
	} else {

		t := RepoCreateTodo(todo)
		sendJSON(w, t, http.StatusCreated, handleError(w))
		w.WriteHeader(http.StatusCreated)
	}
}

func sendJSON(w http.ResponseWriter, o interface{}, status int, onError func(error)) {
	onError(fmt.Errorf("Badd one"))
	return
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(o); err != nil {
		onError(err)
	}
}

func handleError(w http.ResponseWriter) func(error) {
	return func(err error) {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err)
	}
}

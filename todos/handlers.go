package todos

import (
	"encoding/json"
	"fmt"
	"github.com/bernos/go-restapi/application"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

type TodoController struct {
	application.Controller
}

func (c *TodoController) Index(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "Welcome!")
	return nil
}

func (c *TodoController) TodoIndex(w http.ResponseWriter, r *http.Request) error {
	return sendJSON(w, todos, http.StatusOK)
}

func (c *TodoController) TodoShow(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
	return nil
}

func (c *TodoController) TodoCreate(w http.ResponseWriter, r *http.Request) error {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		return err
	}

	if err := r.Body.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		return sendJSON(w, err, 422)
	}

	t := RepoCreateTodo(todo)
	return sendJSON(w, t, http.StatusCreated)
}

func sendJSON(w http.ResponseWriter, o interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(o); err != nil {
		return err
	}
	return nil
}

package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/aleicher/go-static-api/routing"
	"github.com/gorilla/mux"
)

func RegisterTodoRoutes(router *mux.Router) {
	routes := routing.Routes{
		routing.Route{
			Name:        "GetTodos",
			Method:      "GET",
			Pattern:     "/todos",
			HandlerFunc: GetTodos,
		},
		routing.Route{
			Name:        "CreateTodo",
			Method:      "POST",
			Pattern:     "/todos",
			HandlerFunc: CreateTodo,
		},
	}
	routing.AddRoutes(router, routes)
}

type Todo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "./data/todos.json"

var todoMutex = new(sync.Mutex)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todoMutex.Lock()
	defer todoMutex.Unlock()

	// Read the data from the file.
	todoData, err := ioutil.ReadFile(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}
	var todos []Todo
	err = json.Unmarshal(todoData, &todos)
	if err != nil {
		log.Println("error:", err)
	}
	RenderJson(w, todos, 200)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todoMutex.Lock()
	defer todoMutex.Unlock()
	// Stat the file, so we can find its current permissions
	fi, err := os.Stat(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to stat the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	// Read the todos from the file.
	todoData, err := ioutil.ReadFile(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	// Decode the JSON data
	var todos []Todo
	if err := json.Unmarshal(todoData, &todos); err != nil {
		http.Error(w, fmt.Sprintf("Unable to Unmarshal comments from data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	// Add a new todo to the in memory slice of todos
	todoStatus, err := strconv.ParseBool(r.FormValue("done"))

	todos = append(todos, Todo{
		ID:    time.Now().UnixNano() / 1000000,
		Title: r.FormValue("title"),
		Done:  todoStatus,
	},
	)

	// Marshal the todos to indented json.
	todoData, err = json.MarshalIndent(todos, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to marshal todos to json: %s", err), http.StatusInternalServerError)
		return
	}

	// Write out the todos to the file, preserving permissions
	err = ioutil.WriteFile(dataFile, todoData, fi.Mode())
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to write todos to data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	RenderJson(w, todos, 201)

}

package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

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
		}}
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
	log.Println("%s", todos)
	RenderJson(w, todos, 200)

}

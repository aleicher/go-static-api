package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	RegisterPingRoutes(router)
	RegisterTodoRoutes(router)
}

func RenderJson(w http.ResponseWriter, object interface{}, status int) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(status)

	if object != nil {
		if err := json.NewEncoder(w).Encode(object); err != nil {
			panic(err)
		}
	}

}

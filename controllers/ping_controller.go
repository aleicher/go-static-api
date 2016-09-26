package controllers

import (
	"fmt"
	"net/http"

	"github.com/aleicher/go-static-api/routing"
	"github.com/gorilla/mux"
)

func RegisterPingRoutes(router *mux.Router) {
	routes := routing.Routes{
		routing.Route{
			Name:        "Ping",
			Method:      "GET",
			Pattern:     "/ping",
			HandlerFunc: Ping,
		}}
	routing.AddRoutes(router, routes)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong!")
}

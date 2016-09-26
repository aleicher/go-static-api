package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/aleicher/go-static-api/controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Path    string
	Name    string
	Handler http.HandlerFunc
}

type Routes []Route

func AppRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	AddRoutes(r, allRoutes())
	return r
}

func AddRoutes(router *mux.Router, routes Routes) {
	for _, route := range routes {
		AddRoute(router, route)
	}
}

func AddRoute(router *mux.Router, route Route) {
	handler := Log(route.Handler, route.Name)
	router.Methods(route.Method).Path(route.Path).Name(route.Name).Handler(handler)
}

func Log(next http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s %s %s", start, name, r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func allRoutes() Routes {
	var routes = Routes{}
	pingRoute := Route{Name: "Ping", Method: "GET", Path: "/ping", Handler: controllers.Ping}
	routes = append(routes, pingRoute)
	return routes
}

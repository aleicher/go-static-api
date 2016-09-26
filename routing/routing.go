package routing

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func AppRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return r
}

func AddRoutes(router *mux.Router, routes Routes) {
	for _, route := range routes {
		AddRoute(router, route)
	}
}

func AddRoute(router *mux.Router, route Route) {
	handler := Log(route.HandlerFunc, route.Name)
	router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
}

func Log(next http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s %s %s", start, name, r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

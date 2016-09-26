package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/aleicher/go-static-api/controllers"
	"github.com/aleicher/go-static-api/routes"
)

func main() {
	port := flag.Int("port", 3000, "port to run")
	flag.Parse()

	router := routes.AppRouter()
	controllers.RegisterRoutes(router)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(*port),
		Handler: router,
	}
	log.Printf("Running server on port %v.", *port)
	server.ListenAndServe()
}

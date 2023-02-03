package main

import (
	"go-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":3002",
		Handler: router,
	}
	log.Println("Server ejecutandose en puerto 3001")
	log.Println(server.ListenAndServe())

}

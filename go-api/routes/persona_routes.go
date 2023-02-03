package routes

import (
	"go-api/controllers"

	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/find", controllers.Get).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")

}

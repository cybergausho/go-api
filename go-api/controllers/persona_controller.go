package controllers

import (
	"encoding/json"
	"go-api/commons"
	"go-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, req *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close() //al finalizar el getAll ejecuta si o si el close.

	db.Find(&personas)
	json, _ := json.Marshal(personas) //con guion bajo ignora el segundo resultado

	commons.SendReponse(writer, http.StatusOK, json)

}

// metodo que consulta un registro
func Get(writer http.ResponseWriter, req *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(req)["id"]
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound) //notfound depende de la api
	}
}

func Save(writer http.ResponseWriter, req *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(req.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error
	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)
	commons.SendReponse(writer, http.StatusCreated, json)

}

func Delete(writer http.ResponseWriter, req *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(req)["id"]
	db.Find(&persona, id)
	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

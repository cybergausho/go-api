package commons

import "net/http"

func SendReponse(writer http.ResponseWriter, status int, data []byte) {
	writer.Header().Set("content-Type", "applicaction/json") //setea respuesta
	writer.WriteHeader(status)                               //asigna el status
	writer.Write(data)                                       //enviamos la data

}

func SendError(writer http.ResponseWriter, status int) {
	data := []byte(`{}`)
	writer.Header().Set("content-Type", "applicaction/json") //setea respuesta
	writer.WriteHeader(status)
	writer.Write(data)
}

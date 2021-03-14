package main

import (
	"net/http"
	"io"
	"encoding/json"
)

type Persona struct{
	Nombre string
	Apellido string
}

func handlerRaiz( response http.ResponseWriter, request *http.Request){
	io.WriteString( response, "Hola Mundo Api!")
}

func handlerUsuarios( response http.ResponseWriter, request *http.Request){
	io.WriteString(response, "Hola Usuarios!")
}

func handlerPersonas( response http.ResponseWriter, request *http.Request){
	persona := Persona {"German", "Boso"}

	jsResponse, err := json.Marshal(persona)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")

	response.Write(jsResponse)
}
func main() {
	http.HandleFunc("/", handlerRaiz)
	http.HandleFunc("/usuarios", handlerUsuarios)
	http.HandleFunc("/personas", handlerPersonas)
	http.ListenAndServe(":8000", nil)
}
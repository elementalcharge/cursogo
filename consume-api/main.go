package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

type Tarea struct{
	UserId int `json:userId` 
	Id int `json:id`
	Title string `json:title`
	Completed bool `json:completed`
}


func main(){
	var urlApi = "https://jsonplaceholder.typicode.com/todos"


	var cliente = &http.Client{ Timeout: 10 * time.Second}

	response, err := cliente.Get(urlApi)

	if err != nil {
		panic(err.Error())
	}

	var tareas [] Tarea 

	json.NewDecoder(response.Body).Decode(&tareas)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tareas)
}
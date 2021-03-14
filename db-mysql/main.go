package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Persona struct {
	ID int
	Nombre string
	Apellidos string
}

func main() {
	fmt.Println ("Curso de Go. Conexion a MySQL")

	db, err := sql.Open( "mysql", "test:test@tcp(127.0.0.1:3306)/go")

	if err != nil{
		panic (err.Error())
	}

	fmt.Println("Conectado")

	defer db.Close()

	results, err := db.Query("SELECT id, nombre, apellidos FROM persona WHERE nombre = 'German' ")

	if err != nil{
		panic (err.Error())
	}

	for results.Next(){
		var persona Persona

		err=results.Scan(&persona.ID, &persona.Nombre, &persona.Apellidos)

		if err != nil{
			panic(err.Error())
		}

		fmt.Println(persona)
	}
}
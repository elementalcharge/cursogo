package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type Casa struct {
	gorm.Model
	Direccion string
	Numero string
}


func main() {
	fmt.Println("Curso de Go. Conexion a MySQL con ORM")

	db, err := gorm.Open("mysql","test:test@(localhost)/go?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		panic(" Error al conectar a la base de datos")
	}

	defer db.Close()

	fmt.Println("Conectado")

	db.AutoMigrate(&Casa{})

	db.Create(&Casa{Direccion : "Mi direccion", Numero : "25 - 1A"})

	var casa Casa
	var casa2 Casa
	db.First(&casa, 1)

	db.First(&casa2, "direccion = ?", "Mi direccion")

	fmt.Println(casa)

	fmt.Println(casa2)

	db.Model(&casa).Update("Direccion", "Mi nueva direccion")

	db.Delete(&casa)

	var casas [] Casa

	db.Where("numero LIKE?", "%25%").Find(&casas)


	fmt.Println("")
	fmt.Println(casas)




}
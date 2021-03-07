package main

import (
	"fmt"
	"github.com/elementalcharge/cursogo/arrays"
	"github.com/elementalcharge/cursogo/operacionesstrings"
	"github.com/elementalcharge/cursogo/utilidades"
)
const holaMundo string = "Hola %s %s, bienvenido al curso de Go"
func main () {
	name := utilidades.GetName()
	LasName := "Boso"
	var miNumero = utilidades.Suma(40, 60)
	x, y, z := utilidades.GetMultiplesVariables()
	decimal32, decimal64 := utilidades.GetDecimal()

	hola := "Hola"

	
	fmt.Println("hola ", name)
	fmt.Printf(holaMundo, name, LasName)
	fmt.Println("")
	fmt.Println(miNumero, x, y ,z)
	
	fmt.Println(decimal32, decimal64)

	fmt.Println(utilidades.GetUnicode())
	fmt.Println(string(hola[0]))
	fmt.Println(len(hola))
	arrays.ImprimeArray()

	arrays.ImprimirSlice()


	utilidades.Multiplo5()
	utilidades.Multiplo5switch()
	utilidades.Bucles()
	operacionesstrings.OperacionesConStrings()
}
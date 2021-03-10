package main

import (
	"fmt"
	//"github.com/elementalcharge/cursogo/arrays"
	//"github.com/elementalcharge/cursogo/operacionesstrings"
	//"github.com/elementalcharge/cursogo/utilidades"
	//"github.com/elementalcharge/cursogo/maps"
	//"github.com/elementalcharge/cursogo/structs"
	"strconv"
)
//const holaMundo string = "Hola %s %s, bienvenido al curso de Go"


func main () {
	/*name := utilidades.GetName()
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
	*/
	/*
	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetKeyMap("Casa"))
	*/
	/*
	var miVariable structs.Boso = "Mi propio tipo de dato"
	fmt.Println(miVariable)

	antonio := structs.Persona {
		Nombre : "Antonio Jose",
		Apellido : "Galisteo Cantero",
		DocumentoIdentidad : "6699",
		Telefono : []string {"111", "123"},
		Direccion : "YYYY",
		Edad : 30,
	}

	fmt.Println(antonio)

	
	maria := structs.Persona {
		Nombre : "Maria",
		Apellido : "Alvarado",
		DocumentoIdentidad : "YYY",
		Telefono : []string {"222", "345"},
		Direccion : "ZZZZ",
		Edad : 31,
	}

	fmt.Println(maria)

	jorge := new (structs.Persona)
	jorge.Nombre = "Jorge"
	jorge.Apellido = "Apellidos"
	jorge.DocumentoIdentidad = "ZZZ"
	jorge.Telefono = []string {"333", "678"}
	jorge.Direccion = "AAAA"
	jorge.Edad = 31

	fmt.Println(jorge)
	casa := structs.Casa{
		NumeroCasa : 1,
		Personas : []structs.Persona{antonio, maria, *jorge},

	}
	fmt.Println(casa)

	casa.GetNumeroCasa()
	casa.GetPersonasCasa()



*/
/*
	miNumero, error := utilidades.Suma("60",40)
	if error !=nil {
		panic(error)
	}

	fmt.Println(miNumero)
*/
	//punteros()
	canal := make( chan string)
	lanzaHilos(300, canal)

	for valor := range canal{
		fmt.Println(valor)
	}
}
/*
func punteros(){
	x := 100
	var y *int 
	y = &x

	fmt.Println(&x,y)
	*y = 500
	fmt.Println(x,y)
	fmt.Println(x,*y)
	fmt.Println(&x,y)
}*/

func holaMundo( i int , canal chan<- string){
	canal <- "Hola Mundo Numero: " + strconv.Itoa(i)
}

func lanzaHilos( numHilos int, canal chan<- string){
	for i:=0 ; i< numHilos ; i++{
		go holaMundo(i, canal)
		}
	
}
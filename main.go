package main

import (
	"fmt"
	"strings"
)
const holaMundo string = "Hola %s %s, bienvenido al curso de Go"
func main () {
	//name := getName()
	//LasName := "Boso"
	//var miNumero = suma(40, 60)
	//x, y, z := getMultiplesVariables()
	//decimal32, decimal64 := getDecimal()

	//hola := "Hola"

	
	//fmt.Println("hola ", name)
	//fmt.Printf(holaMundo, name, LasName)
	//fmt.Println("")
	//fmt.Println(miNumero, x, y ,z)
	
	//fmt.Println(decimal32, decimal64)

	//fmt.Println(getUnicode())
	//fmt.Println(string(hola[0]))
	//fmt.Println(len(hola))
	//imprimeArray()

	//imprimirSlice()


	//multiplo5()
	//multiplo5switch()
	//bucles()
	operacionesConStrings()
}
func sumoOresto (name string) {
		primerValor  := 0
		segundoValor := 0
		fmt.Println("Dime %s para sumar coloque el primer numero, para restar agregue - a uno de los dos numeros", name)
		fmt.Scanf("%s", &primerValor)
		fmt.Println("Ahora dime %s el segundo numero no te olvides de agregar - si deseas restar y aun no los hecho", name)
		fmt.Scanf("%s", &segundoValor)
		fmt.Println("resultado: %s ",suma( primerValor, segundoValor ) )
		
	}
func getName() string{

	var name string = "nombre por defecto"
	
	fmt.Print("Ingresa tu nombre:  ")

	fmt.Scanf("%s", &name)

	return name

}

func getMultiplesVariables() (int, int, int){
	return 1,2,3
}

func getDecimal() (float32, float64){
	return float32(0.1), float64(float32(0.1))
}

func suma(a int, b int) int{
	return a+b
}
func getUnicode() string {
	return "Goコースへようこそ"
}
func imprimeArray() {
	var array1 [2] string
	array1[0] = "hola"
	array1[1] = "Mundo"
	fmt.Println(array1)
	fmt.Println(len(array1))

	array2 := [4] int {1,2,3,4}
	fmt.Println(array2)
	fmt.Println(len(array2))

	var matriz[2][2] string

	matriz[0][0] = "hola"
	matriz[0][1] = "Mundo"
	matriz[1][0] = "Curso"
	matriz[1][1] = "go"

	fmt.Println(matriz)

	matriz2 := [2][2] int {{1,2}, {3,4}}
	fmt.Println(matriz2)
}

func imprimirSlice(){
	var slice1 [] string
	
	slice1 = append(slice1, "Hola", "slice")
	
	fmt.Println(slice1)

	fmt.Println(len(slice1))

	slice1 = append(slice1, "NuevoElemento")
	fmt.Println(slice1)
	fmt.Println(len(slice1))

	// Matrices

	matriz := [][] string {{"Hola", "Mundo"}, {"curso", "de go"}}

	fmt.Println(matriz)

	var matriz2 [][] string

	row1 := [] string {"Hola2", "SliceMatriz2"}
	row2 := [] string {"curso2", "go2"}

	matriz2 = append(matriz2, row1)
	matriz2 = append(matriz2, row2)

	fmt.Println(matriz2)


}
func getNumero() int{

	var numero = 0

	fmt.Println(" Ingresa un numero")
	fmt.Scanf("%d", &numero)
	fmt.Println("numero es %s", string(numero))
	return numero
}
func multiplo5(){

	if numero := getNumero(); numero % 5 == 0 {
		fmt.Println("Es multiplo de 5")

	} else if numero := getNumero(); numero % 3 == 0 {
		fmt.Println("Es multiplo de 3")

	} else{
		fmt.Println("No es multiplo de 3 o 5")
	}

}

func multiplo5switch(){
	var numero = 0

	fmt.Println(" Ingresa un numero")
	fmt.Scanf("%d", &numero)

	switch numero {
	case 2:
		fmt.Println(" El numero es 2")
	default:
		fmt.Println(" El numero no es 2")
	}

	switch{
	case numero % 5 == 0 :
		fmt.Println("Es multiplo de 5")

	case numero % 3 == 0 :
		fmt.Println("Es multiplo de 3")

	default:
		fmt.Println("No es multiplo de 3 o 5")
	}
}

func bucles (){
	for i:=0; i<3;i++{
		fmt.Println("bucle for ",i)
	}
	j := 50
	for j > 0 {
		j-=10
		fmt.Println("while ",j )
	}
	k := 500
	for {
		k -= 1
		if k == 0{
			fmt.Println("termino bucle infinito")
			break	
		}

	}
}

func operacionesConStrings(){
	var texto = "Hola go, Hola german, Hola mundo"

	var texto2 = "Cadena 2"

	fmt.Println(texto)

	fmt.Println(strings.ToUpper(texto))

	fmt.Println(strings.ToLower(texto))

	fmt.Println(strings.Replace(texto, "Hola", "Adios", -1))

	fmt.Println(strings.Compare(texto, texto2))

	fmt.Println(strings.Split(texto, ","))

	fmt.Println(strings.Contains(texto, "Hola"))

	fmt.Println(strings.Contains(texto, "coche"))



}
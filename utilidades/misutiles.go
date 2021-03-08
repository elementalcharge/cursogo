package utilidades

import (
	"fmt"
	"errors"
)
// SumoOresto intento de calculadora sin uso de condicionales para detectar operacion Go
/*func SumoOresto (name string) {
		primerValor  := 0
		segundoValor := 0
		fmt.Println("Dime %s para sumar coloque el primer numero, para restar agregue - a uno de los dos numeros", name)
		fmt.Scanf("%s", &primerValor)
		fmt.Println("Ahora dime %s el segundo numero no te olvides de agregar - si deseas restar y aun no los hecho", name)
		fmt.Scanf("%s", &segundoValor)
		fmt.Println("resultado: %s o falle por %s",Suma( primerValor, segundoValor ) )
		
	}
*/
// GetName pide un nombre para devolver como string o devuelve "nombre por defecto" en caso de enviar vacio	
func GetName() string{

	var name string = "nombre por defecto"
	
	fmt.Print("Ingresa tu nombre:  ")

	fmt.Scanf("%s", &name)

	return name

}

func GetMultiplesVariables() (int, int, int){
	return 1,2,3
}

func GetDecimal() (float32, float64){
	return float32(0.1), float64(float32(0.1))
}

func Suma(a interface{}, b interface{}) (int, error)  {
	switch a.(type){
	case string: return 0, errors.New("El valor A NO es un entero")
	}
	switch b.(type){
	case string: return 0, errors.New("El valor A NO es un entero")
	}
	return a.(int)+b.(int), nil

}
func GetUnicode() string {
	return "Goコースへようこそ"
}

func GetNumero() int{

	var numero = 0

	fmt.Println(" Ingresa un numero")
	fmt.Scanf("%d", &numero)
	fmt.Println("numero es %s", string(numero))
	return numero
}
func Multiplo5(){

	if numero := GetNumero(); numero % 5 == 0 {
		fmt.Println("Es multiplo de 5")

	} else if numero := GetNumero(); numero % 3 == 0 {
		fmt.Println("Es multiplo de 3")

	} else{
		fmt.Println("No es multiplo de 3 o 5")
	}

}

func Multiplo5switch(){
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

func Bucles (){
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

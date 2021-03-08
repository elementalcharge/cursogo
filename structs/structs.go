package structs

import "fmt"

type Boso string

type Persona struct{
	Nombre string
	Apellido string
	DocumentoIdentidad string
	Telefono []string
	Direccion string
	Edad int
}

type Casa struct {
	NumeroCasa int
	Personas []Persona
}


func (c Casa) GetNumeroCasa() {
	fmt.Println("el numero de la casa es: ", c.NumeroCasa)

}

func (c Casa) GetPersonasCasa() {
	defer SoyDefer()
	fmt.Println("el numero de la casa es: ", c.Personas)

}

func SoyDefer(){
	fmt.Println("Defer, esto quiero que se ejecute al final")
}
package arrays

import "fmt"

// ImprimeArray muestra como se trabaja con arrays y matrices en Go
func ImprimeArray() {

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

// ImprimeArray muestra como se trabaja con Slice en Go
func ImprimirSlice(){
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
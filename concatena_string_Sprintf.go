package main

import "fmt"

func main() {

	cadena_1 := "todo"
	cadena_2 := " tiene"
	cadena_3 := " su final"
	cadena_4 := " feliz"
	cadena_5 := " o no?"

	str_final := fmt.Sprintf("%s%s%s%s%s", cadena_1, cadena_2, cadena_3, cadena_4, cadena_5)
	/* aquí se agrega un "%s" por cada cadena a imprimir */

	fmt.Println(str_final)

}

package main

import "fmt"

func main() {

	fmt.Println(Suma("1", "4", "6"))
	fmt.Println(resta(25, 15, 5))

}

func Suma(numero1, numero2, numero3 string) string {

	return numero1 + numero2 + numero3

}

func resta(tres, dos, uno int) int {

	return tres - dos - uno

}

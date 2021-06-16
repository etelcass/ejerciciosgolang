package main

import "fmt"

func main() {

	str_1 := "Hello "

	str_1 += "World"
	str_1 += " madafaka and"

	str_2 := "hola "

	str_2 += str_1 + " munda (; "

	fmt.Println("Final String:", str_2)

}

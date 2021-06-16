package main

import (
	"bytes"
	"fmt"
)

func main() {
	var byte_buf bytes.Buffer

	byte_buf.WriteString("Hola ")
	byte_buf.WriteString("intentando ")
	byte_buf.WriteString("ver como ")
	byte_buf.WriteString("funcionan ")
	byte_buf.WriteString("todos ")
	byte_buf.WriteString("los ")
	byte_buf.WriteString("espacio")

	fmt.Println(byte_buf.String())
}

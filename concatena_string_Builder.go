package main

import (
	"fmt"
	"strings"
)

func main() {
	var str_build strings.Builder

	str_build.WriteString("Hello ")
	str_build.WriteString("i'm")
	str_build.WriteString(" cheking the ")
	str_build.WriteString("space")
	fmt.Println(str_build.String())
}

package main

import (
	"fmt"
	"strings"
)

func main() {

	str_slices := []string{"esto", "es lo", "que tiene", "que imprimirse"}
	str_concat := strings.Join(str_slices, " ") /* entre las comillas " " puedes poner lo que separará	las cadenas,
	ya sea un espacio, guion, etc. */

	fmt.Println(str_concat)

}

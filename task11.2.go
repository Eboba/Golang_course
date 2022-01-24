package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Функции для работы со строками")

	var a, b, с string
	a = "Golang is programming language"
	b = "a"
	с = ""
	i := strings.ReplaceAll(a, b, с)

	fmt.Println(len(a) - len(i))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Строки и приведения типов")

	fmt.Println("Напишите число:")
	var a string
	fmt.Scan(&a)

	fmt.Println("Укажите систему для перевода:")
	var b int
	fmt.Scan(&b)

	i, err := strconv.ParseInt(a, b, 0)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}

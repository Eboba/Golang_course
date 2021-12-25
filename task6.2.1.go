package main

import (
	"fmt"
)

func main() {

	fmt.Println("Задача 3. Возведение в степень.")

	var num int
	var degree int
	answer := 1

	fmt.Println("Укажите число:")
	fmt.Scan(&num)
	fmt.Println("Укажите степень:")
	fmt.Scan(&degree)

	if degree == 0 {
		fmt.Println(num, "в степени 0 = 1")

	} else if degree < 0 {
		fmt.Println("Укажите степень больше либо равной нулю")

	} else if degree > 0 {
		for i := 0; i != degree; i = i + 1 {
			answer = answer * num
		}
	}

	if degree > 0 {
		fmt.Println(num, "в степени", degree, "=", answer)
	}
}

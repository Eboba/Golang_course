package main

import (
	"fmt"
)

func main() {

	fmt.Println("Складывание всех чисел кратных 4")

	var sum int
	var num int

	for num <= 40 {
		num = num + 1

		if num%4 != 0 {
			continue
		}
		sum = sum + 4
		fmt.Println(sum)
	}

}

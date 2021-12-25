package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проверка пароля.")

	fmt.Println("Укажите пароль:")
	var answer string

	for {
		fmt.Scan(&answer)
		if answer == "admin" {
			break
		} else {
			fmt.Println("Пароль не верный")
			fmt.Println("Укажите пароль:")
		}
	}
	fmt.Println("Пароль верный")
}

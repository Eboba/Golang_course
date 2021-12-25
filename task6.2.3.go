package main

import (
	"fmt"
)

func main() {

	fmt.Println("Часовой")

	var coordinate int
	path := 1

	for {
		coordinate = path + coordinate

		if coordinate == 10 {
			path = -1

		} else if coordinate == -10 {
			path = 1
		}

		if coordinate == 0 {
			fmt.Println("Часовой прошел КПП!")
		} else {
			fmt.Println(coordinate)
		}
	}

}

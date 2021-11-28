package main

import (
	"fmt"
)

func main() {

	var bamboo int = 100
	var varBamboo int = 50
	var pests int = 20

	fmt.Println("Высоты бамбука в середине третьего дня", bamboo+varBamboo*25/10-pests*2)
}

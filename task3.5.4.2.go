package main

import (
	"fmt"
)

func main() {

	var bamboo int = 100
	var varBamboo int = 50
	var pests int = 20
	var limitBamboo int = 3000

	fmt.Println("Чтобы можно было срубить и продать бамбуку понадобится ", (limitBamboo-bamboo)/(varBamboo-pests), "полных дней")
}

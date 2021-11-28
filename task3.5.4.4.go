package main

import (
	"fmt"
)

func main() {

	var bamboo int = 0
	var varBamboo int = 0
	var pests int = 0
	var limBamboo int = 0

	fmt.Println("Введите длину саженца(см):")
	fmt.Scan(&bamboo)

	fmt.Println("Введите скорость роста бамбука(см):")
	fmt.Scan(&varBamboo)

	fmt.Println("Введите скорость поедания гусеницами(см):")
	fmt.Scan(&pests)

	fmt.Println("Введите необходимую высоту бамбука(см):")
	fmt.Scan(&limBamboo)

	fmt.Println("Чтобы можно было срубить и продать бамбуку понадобится ", (limBamboo-bamboo)/(varBamboo-pests), "полных дней")
}

package main

import (
	"fmt"
)

func main() {

	var bamboo int = 0
	var varBamboo int = 0
	var pests int = 0
	var day int = 0

	fmt.Println("Введите длину саженца(см):")
	fmt.Scan(&bamboo)

	fmt.Println("Введите скорость роста бамбука(см):")
	fmt.Scan(&varBamboo)

	fmt.Println("Введите скорость поедания гусеницами(см):")
	fmt.Scan(&pests)

	fmt.Println("Введите количество дней:")
	fmt.Scan(&day)

	fmt.Println("Высоты бамбука в конце", day, "дня составит:", day*(varBamboo-pests)+bamboo+varBamboo, "см.")
}

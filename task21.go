package main

import (
	"fmt"
)

func task21() {
	var CircleNumber int = 4
	var Speed int = 358
	var Engine int = 254
	var Wheels int = 93
	var Eudder int = 49
	var Wing int = 21
	var Rain int = 14

	fmt.Println("==============")
	fmt.Println("Супер Гонки", "Круг", CircleNumber)
	fmt.Println("==============")
	fmt.Print("Шумахер (", Speed, ")", "\n")
	fmt.Println("==============")
	fmt.Println("Водитель: Шумахер", "скорость:", Speed)
	fmt.Println("-------------")
	fmt.Println("Оснащение")
	fmt.Println("Двигатель:", Engine)
	fmt.Println("Колеса:", Wheels)
	fmt.Println("Руль:", Eudder)
	fmt.Println("-------------")
	fmt.Println("Действия плохой погоды")
	fmt.Println("Ветер:", Wing)
	fmt.Println("Дождь:", Rain)
}

package main

import (
	"fmt"
)

func main() {

	fmt.Println("-Как твое имя сынок?")
	var name string
	fmt.Scan(&name)
	fmt.Println("-Меня зовут", name, "сэр!")

	fmt.Print("-Откуда ты родом, ", name, " ?\n")
	var planet string
	fmt.Scan(&planet)
	fmt.Print("-Из ", planet, ", сэр! Планеты проституток и хоккеистов!\n")

	fmt.Print("-Моя жена тоже с ", planet, ".\n")
	fmt.Println("-И за какую звездную систему она играет, сэр?")
	var star string
	fmt.Scan(&star)

	fmt.Print("-Сколько ты хочешь за свое тупое чувство юмора?\n")
	var money int
	fmt.Scan(&money)
	fmt.Println("-Я хочу", money, "кредитов.")
	fmt.Println("")
	fmt.Print("«Как вам, ", name, " известно, мы раса мирная, поэтому на наших военных кораблях используются наёмники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты ", planet, " системы ", star, ".\n")
	fmt.Println("")
	fmt.Print("Но случилась неприятность: в связи с большими потерями в последнее время престиж профессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар планеты ", planet, ", впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!\n")
	fmt.Println("")
	fmt.Print(name, ", Вы должны прибыть на планету ", planet, " системы ", star, " и помочь выполнить план призыва. Мы готовы выплатить вам премию в ", money, " кредитов за эту маленькую услугу».\n")
}

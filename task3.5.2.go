package main

import (
	"fmt"
)

func main() {

	var people int = 0
	var varPeople int = 0
	var money int = 0

	fmt.Println("Остановка «Улица Программистов».")
	fmt.Println("В салоне пассажиров:", people)
	fmt.Println("Сколько пассажиров вошло на остановке:")
	fmt.Scan(&people)
	money += people * 20

	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку «Проспект Алгоритмов».")
	fmt.Println("В салоне пассажиров:", people)
	fmt.Println("Сколько пассажиров вошло на остановке:")
	varPeople = 0
	fmt.Scan(&varPeople)
	people += varPeople
	money += varPeople * 20
	fmt.Println("Сколько пассажиров вышло на остановке:")
	fmt.Scan(&varPeople)
	people -= varPeople

	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку «Аллея функций».")
	fmt.Println("В салоне пассажиров:", people)
	fmt.Println("Сколько пассажиров вошло на остановке:")
	varPeople = 0
	fmt.Scan(&varPeople)
	people += varPeople
	money += varPeople * 20
	fmt.Println("Сколько пассажиров вышло на остановке:")
	fmt.Scan(&varPeople)
	people -= varPeople

	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на конечную остановку «Бульвар дедлайна».")
	fmt.Println("В салоне пассажиров:", people)

	fmt.Println("Всего заработали:", money, "рублей.")
	fmt.Println("Зарплата водителя:", money/4, "рублей.")
	fmt.Println("Расходы на топливо:", money/5, "рублей.")
	fmt.Println("Налоги:", money/5, "рублей.")
	fmt.Println("Расходы на ремонт машины:", money/5, "рублей.")
	fmt.Println("Итого доход:", money-(money*6/10)-(money/4), "рублей.")
}

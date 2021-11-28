package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите логин:")
  var login string
  fmt.Scan (&login)

  fmt.Println("Введите пароль:")
  var pass string
  fmt.Scan (&pass)

  fmt.Println("Укажите возраст:")
  var age int
  fmt.Scan (&age)

  fmt.Println("Поздравляю,", login, ", теперь вы зарегистрированы!")
  fmt.Println("Ваш пароль:", pass)
  fmt.Println("Ваш возраст:", age)
}
}


===================================

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите логин:")
  var login string
  fmt.Scan (&login)

  fmt.Println("Введите пароль:")
  var pass string
  fmt.Scan (&pass)

  fmt.Println("-----")
  fmt.Println(login, ", вы успешно зашли!")
}

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Калькулятор стоимости товара со скидкой.")
	fmt.Println("Стоимость товара:")
  var priceItemRub int
  fmt.Scan(&priceItemRub)
	fmt.Println("Стоимость доставки:")
  var DeliveryRub int
  fmt.Scan(&DeliveryRub)
	fmt.Println("Размер скидки:")
  var DiscountRub int 
  fmt.Scan(&DiscountRub)
	fmt.Println("---------")
	fmt.Println("Итого:", priceItemRub + DeliveryRub - DiscountRub)

}



=============================================


package main

import (
	"fmt"
)

func main() {

	fmt.Println("квадрат числа")
  
  var a int
  fmt.Scan(&a)

	fmt.Println("равен:", a*a)

}
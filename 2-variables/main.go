package main

import (
	"fmt"
)

const (
	appleCost float32 = 5.99
	pearCost float32= 7.00
	budget float32= 23.00
)

func main() {
	fmt.Printf("Маємо %v грн\n", budget)
	fmt.Printf("Ціна яблука %v грн та груші %v грн\n", appleCost, pearCost)
	fmt.Printf("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n")
	var sum = (appleCost*9+pearCost*8)
	fmt.Printf("Потрібна сума: %v грн\n", sum)
	fmt.Printf("Скільки груш ми можемо купити?\n")
	var pearToBuy = (budget / pearCost)
	fmt.Printf("За свої кошти можемо придбати: %v груші\n", int(pearToBuy))
	fmt.Printf("Скільки яблук ми можемо купити?\n")
	var appleToBuy = (budget / appleCost)
	fmt.Printf("За свої кошти можемо придбати: %v яблука\n", int(appleToBuy))
	fmt.Printf("Чи ми можемо купити 2 груші та 2 яблука?\n")
	var check = (pearCost*2+appleCost*2)
	fmt.Printf("ВІдповідь: %v \n", budget >= check)
	

}

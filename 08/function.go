package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//Oddiy function'dan foydalanish
	sum := sum(34, 23)
	fmt.Println("sum:", sum)

	//Bir nechta qiymat qaytaruvhi function'dan foydalanish
	sqrtValue, error := sqrt(92)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(sqrtValue)
	}

	//Function'ga argument berishni birinchi yo'qli
	a, b := 10, 12
	fmt.Println("\nFunction'ga argument berishni birinchi yo'qli")
	fmt.Println("a:", a, "b:", b)
	swap(a, b)
	fmt.Println("a:", a, "b:", b)

	//Function'ga argument berishni ikkinchi yo'li
	fmt.Println("\nFunction'ga argument berishni ikkinchi yo'li")
	fmt.Println("a:", a, "b:", b)
	swapByReference(&a, &b)
	fmt.Println("a:", a, "b:", b)
}

//Oddiy function
func sum(a int, b int) int {
	return a + b
}

//Bir nechta qiymat qaytaruvchi function
func sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("Argument sifatida manfiy qiymat berilmoqda")
	}
	return math.Sqrt(number), nil
}

func swap(a int, b int) {
	o := a
	a = b
	b = o
}

func swapByReference(a, b *int) {
	o := *a
	*a = *b
	*b = o
}

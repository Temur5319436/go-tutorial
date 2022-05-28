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
	}
	fmt.Println(sqrtValue)
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

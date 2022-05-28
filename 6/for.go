package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//For loop'ni ishlatilishi.
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Fordan while ko'rinishida ishlatish
	i := 0
	for i < 4 {
		fmt.Println("For as while : ", i)
		i++
	}

	//For loop'ni infinity ko'rinishida ishlatish
	for {
		fmt.Println("For " + strconv.Itoa(time.Now().Second()))
		time.Sleep(1 * time.Second)
		break
	}

	//For loop'ni array bilan qo'lanishi
	myArray := [3]string{
		"Temur",
		"Usmonaliyev",
		"Baxtiyorjon o'gli",
	}

	for index, value := range myArray {
		fmt.Println("array index :", index, "value :", value)
	}

	//For loop'ni map bilan qo'lanishi
	myMap := map[int]string{
		200: "ok",
		400: "Invalid request",
		404: "Not found",
		500: "Internal server error",
	}
	for key, value := range myMap {
		fmt.Println("map key:", key, "value:", value)
	}
}

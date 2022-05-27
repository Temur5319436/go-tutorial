package main

import "fmt"

func main() {
	var cars [3]string

	cars[0] = "Nexia"
	cars[1] = "Civic"
	cars[2] = "Accord"

	fmt.Println("cars : ")
	fmt.Println("1 : ", cars[0])
	fmt.Println("2 : ", cars[1])
	fmt.Println("3 : ", cars[2])

	var fruits = [3]string{"Apple", "Orange", "Banana"}

	fmt.Println("\nfruits : ")
	fmt.Println("1 : ", fruits[0])
	fmt.Println("2 : ", fruits[1])
	fmt.Println("3 : ", fruits[2])

}

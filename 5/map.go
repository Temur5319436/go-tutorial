package main

import "fmt"

func main() {
	//Mapni e'lon qilish
	statuses := make(map[string]int)

	statuses["added"] = 1
	statuses["progress"] = 2
	statuses["done"] = 3

	fmt.Println("statuses : ", statuses)

	//Mapdan qiymatni olish
	progressStatus := statuses["progress"]
	fmt.Println("progress : ", progressStatus)

	//Mapdan elementni o'chirib tashlash
	delete(statuses, "progress")
	fmt.Println("statuses : ", statuses)
}

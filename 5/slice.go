package main

import (
	"fmt"
	"sort"
)

func main() {
	//Yangi slice e'lon qilish
	numbers := []int{23, 12, 56, 23, 21, 65, 76, 2}
	fmt.Println("numbers : ", numbers)

	//Slicega yangi element qo'shish
	numbers = append(numbers, 5)

	fmt.Println("numbers : ", numbers)
	fmt.Println("numbers length: ", len(numbers))

	//Arraydan scliceni ajratib olish
	newArray := []int{9, 3, 2, 4, 5, 7, 8}
	newSlice := newArray[1:4]
	fmt.Println("Arraydan ajratib olingan slice: ", newSlice)

	//Sliceni tartiblash
	mySlice := []int{9, 3, 2, 45, 67, 1, 4}
	sort.Ints(mySlice)
	fmt.Println("Tartiblangan slice : ", mySlice)
}

package main

import "fmt"

func main() {
	//O'lchami aniq bo'lgan arraylarni e'lon qiilish
	var cars [3]string
	cars[0] = "Nexia"
	cars[1] = "Civic"
	cars[2] = "Accord"
	fmt.Println("cars : ")
	fmt.Println("1 : ", cars[0])
	fmt.Println("2 : ", cars[1])
	fmt.Println("3 : ", cars[2])

	//O'lchami aniq bo'lgan arraylarni e'lon qiilish vaqtida qiymat qo'shish
	var fruits = [3]string{"Apple", "Orange", "Banana"}
	fmt.Println("\nfruits : ")
	fmt.Println("1 : ", fruits[0])
	fmt.Println("2 : ", fruits[1])
	fmt.Println("3 : ", fruits[2])

	//O'lchami aniq bo'lmagan arraylarni e'lon qiilish vaqtida qiymat qo'shish
	numbers := [...]int{3, 5, 1, 2, 5, 7, 8}
	fmt.Println("numbers: ", numbers)

	//Ko'p o'lchamli arrayni e'lon qiilish.
	people := [3][2]string{
		{"Temur", "Farruh"},
		{"Ulug'bek", "Javlon"},
		{"Bekzod", "Abduqosim"},
	}
	fmt.Println("people: ", people)

	//Arraydan nusxa olishning birinchi yo'li.
	//Bu misolda arraylar xotiradan aloxida joy egallaydi.
	firstArray := [...]string{"Alijon", "Valijon", "Shoxrus"}
	fmt.Println("\nCopy array first way:\nFirst array : ", firstArray)
	copyArray := firstArray
	fmt.Println("Copy array : ", copyArray)
	copyArray[0] = "Ilhomjon"
	fmt.Println("Result first array : ", firstArray)
	fmt.Println("Result second array: ", copyArray)

	//Arraydan nusxa olishning ikkinchi yo'li.
	//Bu misolda copy array birinchi arrayning xotiradagi joyiga link bo'lib qoladi.
	secondArray := [...]string{"Alijon", "Valijon", "Shoxrus"}
	fmt.Println("\nCopy array second way:\nFirst array : ", secondArray)
	secondCopyArray := &secondArray
	fmt.Println("Second way copy from array : ", secondCopyArray)
	secondCopyArray[0] = "Nozimjon"
	fmt.Println("Result first array : ", secondArray)
	fmt.Println("Result second array : ", *secondCopyArray)

}

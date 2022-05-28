package main

import "fmt"

func main() {
	//O'zgaruvchini e'lon qilish
	var a int = 10
	var b int = 2
	fmt.Println("a + b = ", a+b)

	//Bir nechta o'zgaruvchilari e'lon qilish
	var c, d int = 10, 2
	fmt.Println("c - d = ", c-d)

	//O'zgaruvchilni qisqa ko'rinishda e'lom qilish
	e := 10
	f := 2
	fmt.Println("e * f = ", e*f)

	//Bir nechta o'zgaruvchilari qisqa ko'rinishda e'lom qilish
	g, h := 10, 2
	fmt.Println("g / h = ", g/h)

	//Constant'larni e'lon qilish
	const PI = 3.14
	fmt.Println("PI:", PI)
}

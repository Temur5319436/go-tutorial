package main

import (
	"fmt"
	"time"
)

func main() {
	go print("Hello first goroutine !")
	go print("Hello second goroutine !")

	fmt.Scanln()
}

func print(message string) {
	for i := 1; true; i++ {
		fmt.Println(i, message)
		time.Sleep(1 * time.Second)
	}
}

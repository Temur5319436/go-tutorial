package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)

	go randomNumbers(channel)

	for value := range channel {
		fmt.Println(value)
	}
}

func randomNumbers(channel chan int) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 3; i++ {
		randomNumber := rand.Intn(1000)
		time.Sleep(1 * time.Second)
		channel <- randomNumber
	}

	close(channel)
}

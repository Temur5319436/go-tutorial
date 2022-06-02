package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "First message !"
		channel <- "Second message !"
	}()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

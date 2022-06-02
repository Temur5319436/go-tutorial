package main

import (
	"fmt"
	"time"
)

func main() {
	firstChannel := make(chan string)
	secondChannel := make(chan string)

	go func() {
		for {
			firstChannel <- "First channel !"
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			secondChannel <- "Second channel !"
			time.Sleep(600 * time.Millisecond)
		}
	}()

	for {
		select {
		case firstMessage := <-firstChannel:
			fmt.Println(firstMessage)
		case secondMessage := <-secondChannel:
			fmt.Println(secondMessage)
		default:
			fmt.Println("App is waiting ...")
		}
	}
}

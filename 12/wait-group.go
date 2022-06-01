package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		print("Hello first wait group !")
		wg.Done()
	}()
	go func() {
		print("Hello second wait group !")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("App completed !")
}

func print(message string) {
	for i := 1; i <= 7; i++ {
		fmt.Println(i, message)
		time.Sleep(1 * time.Second)
	}
}

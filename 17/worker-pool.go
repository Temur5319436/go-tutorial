package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	jobsCount := 45
	jobs := make(chan calculation, jobsCount)
	results := make(chan calculation, jobsCount)

	for i := 1; i <= 4; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= jobsCount; i++ {
		job := calculation{result: 0, number: i}
		jobs <- job
	}
	close(jobs)

	for i := 1; i < jobsCount; i++ {
		result := <-results
		fmt.Println("FIbonacci (", result.number, "):", result.result)
	}
	close(results)

	fmt.Println("Work has done !\nIt has done", time.Since(start).Milliseconds(), "milliseconds")
}

type calculation struct {
	number, result int
}

func worker(id int, jobs <-chan calculation, results chan<- calculation) {
	for job := range jobs {
		job.result = fibonacci(job.number)
		fmt.Println(id, "worker has done work !")
		results <- job
	}
}

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}

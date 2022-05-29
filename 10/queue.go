package main

import (
	"errors"
	"fmt"
)

func main() {
	q := NewQueue()
	fmt.Println(q)

	fmt.Println("is empty :", q.IsEmpty())

	q.push(3).push(8).push(2)
	fmt.Println("is empty :", q.IsEmpty())

}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) push(x int) *Queue {
	q.data = append(q.data, x)
	return q
}

func (q *Queue) next() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Navbat bo'sh")
	}
	return q.data[0], nil
}

func (q *Queue) shift() bool {
	return true
}

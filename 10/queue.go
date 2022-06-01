package main

import (
	"errors"
	"fmt"
)

func main() {
	q := NewQueue()

	q.push(64)
	fmt.Println("queue elements:", q)
	fmt.Println("is empty :", q.IsEmpty())

	var next, _ = q.next()
	fmt.Println("next element:", next)

	q.push(3).push(8).push(2)
	fmt.Println("queue elements:", q)
	fmt.Println("is empty :", q.IsEmpty())

	next, _ = q.shift()
	fmt.Println("shifted element:", next)
	fmt.Println("queue elements:", q)
	fmt.Println("queue is empty:", q.IsEmpty())
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

func (q *Queue) next() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty !")
	}
	return q.data[0], nil
}

func (q *Queue) push(x int) *Queue {
	q.data = append(q.data, x)
	return q
}

func (q *Queue) shift() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty !")
	}
	x := q.data[0]
	q.data = q.data[1:]
	return x, nil
}

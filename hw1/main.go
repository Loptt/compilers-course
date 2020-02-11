package main

import "fmt"

func main() {
	var q Queue
	q.Push(1)
	q.Push(3)
	q.Push(2)
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
	q.Push(5)
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
	fmt.Println("All queue")

	for !q.Empty() {
		fmt.Println(q.Front())
		q.Pop()
	}
}

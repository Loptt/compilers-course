package main

import (
	"fmt"
	"./hashtable"
	"./queue"
)

func main() {
	// Queue
	var q queue.Queue
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

	// HashTable
	var ht hashtable.HashTable
	ht.Insert(4, "Hola")
	ht.Insert(14, "como")
	ht.Insert(34, "estas")
	ht.Insert(44, "bien")
	fmt.Println(ht.Get(4))
	fmt.Println(ht.Get(14))
	fmt.Println(ht.Get(34))
	fmt.Println(ht.Get(44))
	ht.Remove(4)
	fmt.Println(ht.Get(4))

}

package main

import "fmt"

func main() {
	var s Stack
	s.Push(1)
	s.Push(3)
	s.Push(2)
	fmt.Println(s.Top())
}

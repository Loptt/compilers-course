package main

import "errors"

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	head *Node
}

func (s Stack) Pop() error {
	if s.head == nil {
		return errors.New("Empty stack")
	}

	s.head = s.head.next
	return nil
}

func (s Stack) Top() (int, error) {
	if s.head == nil {
		return 0, errors.New("Empty stack")
	}
	return s.head.val, nil
}

func (s Stack) Push(val int) error {
	newHead := new(Node)
	newHead.val = val
	newHead.next = s.head
	s.head = newHead
	return nil
}

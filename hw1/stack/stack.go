package stack

import (
	"fmt"
	"strings"
)

// Node is a data container
type Node struct {
	val  interface{}
	next *Node
}

// Stack implements the stack data structure
type Stack struct {
	head *Node
}

// String represents the stack structure in a string format
func (s Stack) String() string {
	if s.Empty() {
		return "<Empty Stack>"
	}

	curr := s.head
	var result strings.Builder

	result.WriteString("< ")

	for curr != nil {
		value := fmt.Sprintf("%v", curr.val)
		if curr.next == nil {
			result.WriteString(value)
		} else {
			result.WriteString(value + ", ")
		}
		curr = curr.next
	}

	result.WriteString(" >")

	return result.String()
}

// Equal overrides the stack comparator
func (s Stack) Equal(s2 Stack) bool {
	if s.Empty() && s2.Empty() {
		return true
	}

	curr1 := s.head
	curr2 := s2.head

	for curr1 != nil && curr2 != nil {
		if curr1.val != curr2.val {
			return false
		}
		curr1 = curr1.next
		curr2 = curr2.next
	}

	if curr1 != curr2 {
		return false
	}

	return true
}

// Empty returns true if stack is empty
func (s *Stack) Empty() bool {
	return s.head == nil
}

// Pop removes the first element in the container
func (s *Stack) Pop() {
	if s.Empty() {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *Stack) Top() interface{} {
	if s.Empty() {
		return 0
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *Stack) Push(val interface{}) {
	newHead := &Node{val, s.head}
	s.head = newHead
}

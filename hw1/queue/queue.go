package queue

import (
	"fmt"
	"strings"
)

// Node is a data container
type Node struct {
	val  interface{}
	next *Node
}

// Queue implemenets the queue data structure
type Queue struct {
	head *Node
	tail *Node
}

// Equal overrides the queue comparator
func (q Queue) Equal(q2 Queue) bool {
	if q.Empty() && q2.Empty() {
		return true
	}

	curr1 := q.head
	curr2 := q2.head

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

// String represents the stack structure in a string format
func (q Queue) String() string {
	if q.Empty() {
		return "<Empty Stack>"
	}

	curr := q.head
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

// Empty returns true if container is empty
func (q *Queue) Empty() bool {
	return q.head == nil
}

// Pop removes the first element in the queue
func (q *Queue) Pop() {
	if q.Empty() {
		return
	}

	q.head = q.head.next
}

// Front returns the first element in the queue
func (q *Queue) Front() interface{} {
	if q.Empty() {
		return 0
	}

	return q.head.val
}

// Push adds a new element to the tail
func (q *Queue) Push(val interface{}) {
	newNode := &Node{val: val}
	newNode.val = val

	if q.Empty(){
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

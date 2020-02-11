package main

// Node is a data container
type Node struct {
	val  int
	next *Node
}

// Queue implemenets the queue data structure
type Queue struct {
	head *Node
	tail *Node
}

// Empty returns true if container is empty
func (q *Queue) Empty() bool {
	return q.head == nil
}

// Pop removes the first element in the queue
func (q *Queue) Pop() {
	if q.head == nil {
		return
	}

	q.head = q.head.next
}

// Front returns the first element in the queue
func (q *Queue) Front() int {
	if q.head == nil {
		return 0
	}

	return q.head.val
}

// Push adds a new element to the tail
func (q *Queue) Push(val int) {
	newNode := new(Node)
	newNode.val = val

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

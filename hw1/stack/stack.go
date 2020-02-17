package stack

// Node is a data container
type Node struct {
	val  int
	next *Node
}

// Stack implements the stack data structure
type Stack struct {
	head *Node
}

// Empty returns true if stack is empty
func (s *Stack) Empty() bool {
	return s.head == nil
}

// Pop removes the first element in the container
func (s *Stack) Pop() {
	if s.head == nil {
		return
	}

	s.head = s.head.next
}

// Top returns the first element in the container
func (s *Stack) Top() int {
	if s.head == nil {
		return 0
	}
	return s.head.val
}

// Push adds an element to the top of the container
func (s *Stack) Push(val int) {
	newHead := &Node{val, s.head}
	s.head = newHead
}

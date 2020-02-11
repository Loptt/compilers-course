package main

// TableEntry specifies a field in the hashtable
type TableEntry struct {
	key   int
	value string
	next  *TableEntry
}

// SIZE of the HashTable array implementation
const SIZE = 10

// HashTable implemenets a hash table data structure
type HashTable struct {
	array [SIZE]*TableEntry
}

// Insert adds an element to the table
func (h *HashTable) Insert(key int, value string) {
	index := key % SIZE

	if h.array[index] != nil {
		curr := h.array[index]

		for curr.next != nil {
			if curr.key == key {
				return
			}
			curr = curr.next
		}

		if curr.key == key {
			return
		}

		curr.next = &TableEntry{key, value, nil}
	} else {
		h.array[index] = &TableEntry{key, value, nil}
	}
}

// Get retreives the value corresponding to the key
func (h *HashTable) Get(key int) string {
	index := key % SIZE

	if h.array[index] != nil {
		curr := h.array[index]

		for curr != nil {
			if curr.key == key {
				return curr.value
			}
			curr = curr.next
		}
	}

	return ""
}

// Remove deletes the TableEntry corresponding to the key
func (h *HashTable) Remove(key int) {
	index := key % SIZE

	if h.array[index] != nil {
		curr := h.array[index]

		if curr.key == key {
			h.array[index] = curr.next
			return
		}

		for curr.next != nil {
			if curr.next.key == key {
				curr.next = curr.next.next
				return
			}
			curr = curr.next
		}
	}
}

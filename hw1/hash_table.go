package main

type TableEntry struct {
	key   int
	value string
	next  *TableEntry
}

type HashTable struct {
	array [100]*TableEntry
}

func (h *HashTable) Insert(key int, value string) {
	index := key % 100

	if h.array[index] != nil {
		curr := h.array[index]

		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &TableEntry{key, value, nil}
	} else {
		h.array[index] = &TableEntry{key, value, nil}
	}
}

func (h *HashTable) Get(key int) string {
	index := key % 100

	if h.array[index] == nil {
		return ""
	} else {

	}
}

package main

import "regexp"

// Node represents a node in the linked list
type Node struct {
	key   string
	value int
	next  *Node
}

// HashTable represents the hash table
type HashTable struct {
	size  int
	table []*Node
}

// NewHashTable creates and returns a new HashTable
func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]*Node, size),
	}
}

// HashFunc calculates the hash value for a given key
func (h *HashTable) HashFunc(key string) int {
	var hash int
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % h.size
}

// Insert adds a new key-value pair to the hash table
func (h *HashTable) Insert(key string, value int) {
	hash := h.HashFunc(key)
	newNode := &Node{key: key, value: value}
	if h.table[hash] == nil {
		h.table[hash] = newNode
	} else {
		currentNode := h.table[hash]
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

// Search finds and returns the value for a given key
func (h *HashTable) Search(key string) (int, bool) {
	hash := h.HashFunc(key)
	if h.table[hash] != nil {
		currentNode := h.table[hash]
		for currentNode != nil {
			if currentNode.key == key {
				return currentNode.value, true
			}
			currentNode = currentNode.next
		}
	}
	return 0, false
}

func (h *HashTable) SearchRegex(pattern string) ([]int, bool) {
	var result []int
	var found bool
	for _, currentNode := range h.table {
		for currentNode != nil {
			matched, _ := regexp.MatchString(pattern, currentNode.key)
			if matched {
				result = append(result, currentNode.value)
				found = true
			}
			currentNode = currentNode.next
		}
	}
	return result, found
}

/*
The main advantage of using linked lists to handle collisions is that it allows the hash table to continue working even if the hash function is not perfect and generates collisions. It also allows for the storage of multiple items with the same hash value. However, there are also some drawbacks to chaining:

1. As the number of items in the hash table increases, the length of the linked lists can become very large, and searching for an item in a large linked list can be slow.

2. Chaining can also use a lot of memory, since each key-value pair is stored in a separate node in the linked list.
*/

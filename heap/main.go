package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"strings"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Peek() int          { return h[0] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	var query, choice, element int
	fmt.Scan(&query)
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < query; i++ {
		fmt.Scan(&choice, &element)
		if choice == 1 {
			heap.Push(h, element)
		} else if choice == 2 {
			heap.Remove(h, element)
		} else {
			fmt.Println(h.Peek())
			fmt.Println(h)
		}
	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

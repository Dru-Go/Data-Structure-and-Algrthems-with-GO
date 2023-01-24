package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumMoves' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING_ARRAY grid
 *  2. INTEGER startX
 *  3. INTEGER startY
 *  4. INTEGER goalX
 *  5. INTEGER goalY
 */

type Cell struct {
	i int32
	j int32
}

// Create Q node
type QNode struct {
	data Cell
	next *QNode
}

func getQNode(value Cell) *QNode {
	// return new QNode
	return &QNode{
		value,
		nil,
	}
}

type MyQueue struct {
	head  *QNode
	tail  *QNode
	count int
}

func getMyQueue() *MyQueue {
	// return new MyQueue
	return &MyQueue{
		nil,
		nil,
		0,
	}
}
func (this MyQueue) size() int {
	return this.count
}
func (this MyQueue) isEmpty() bool {
	return this.count == 0
}

// Add new node of queue
func (this *MyQueue) enqueue(value Cell) {
	// Create a new node
	var node *QNode = getQNode(value)
	if this.head == nil {
		// Add first element into queue
		this.head = node
	} else {
		// Add node at the end using tail
		this.tail.next = node
	}
	this.count++
	this.tail = node
}

// Delete a element into queue
func (this *MyQueue) dequeue() Cell {
	if this.head == nil {
		fmt.Println("Empty Queue")
		return -1
	}
	// Pointer variable which are storing
	// the address of deleted node
	var temp *QNode = this.head
	// Visit next node
	this.head = this.head.next
	this.count--
	if this.head == nil {
		// When deleting a last node of linked list
		this.tail = nil
	}
	return temp.data
}

// Get front node
func (this MyQueue) peek() Cell {
	if this.head == nil {
		fmt.Println("Empty Queue")
		return Cell{}
	}
	return this.head.data
}

func bft(array [][]rune, dist [][]int32, startX int32, startY int32, goalX int32, goalY int32) {
	dist[startX][startY] = 0
	var start Cell = Cell{startX, startY}
	queue := getMyQueue()
	queue.enqueue(start)

	for !queue.isEmpty() {

	}
}

func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	// Write your code here

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	startXTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	startX := int32(startXTemp)

	startYTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	startY := int32(startYTemp)

	goalXTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	goalX := int32(goalXTemp)

	goalYTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	goalY := int32(goalYTemp)

	result := minimumMoves(grid, startX, startY, goalX, goalY)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// Return the number of items in the stack
func (this *Stack) isEmpty() bool {
	return this.length == 0
}

// View the top item on the stack
func (this *Stack) peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func shiftElement(stack1, stack2 *Stack) {
	for !stack1.isEmpty() {
		val := stack1.pop()
		stack2.push(val)
	}
	// fmt.Println("Stack 1", stack1)
	// fmt.Println("Stack 2", stack2)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var query, choice, element int

	reader := bufio.NewReader(os.Stdin)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	query = int(nTemp)
	var stack1 = New()
	var stack2 = New()
	for i := 0; i < query; i++ {
		reader := bufio.NewReader(os.Stdin)
		str, _, _ := reader.ReadLine()
		split := strings.Split(string(str), " ")
		nTemp1, _ := strconv.ParseInt(split[0], 10, 64)
		if len(split) > 1 {
			nTemp2, _ := strconv.ParseInt(split[1], 10, 64)
			element = int(nTemp2)
		}
		choice = int(nTemp1)
		// fmt.Println(choice)
		// fmt.Println(element)
		if choice == 1 {
			stack1.push(element)
		} else if choice == 2 {
			if stack2.isEmpty() {
				shiftElement(stack1, stack2)
			}
			stack2.pop()
		} else {
			if stack2.isEmpty() {
				shiftElement(stack1, stack2)
			}
			val := stack2.pop()
			fmt.Println(val)
		}
	}

	fmt.Println(0)
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

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
func (ss *Stack) Len() int {
	return ss.length
}

// Return the number of items in the stack
func (ss *Stack) isEmpty() bool {
	return ss.length == 0
}

// View the top item on the stack
func (ss *Stack) peek() interface{} {
	if ss.length == 0 {
		return nil
	}
	return ss.top.value
}

// Pop the top item of the stack and return it
func (ss *Stack) pop() interface{} {
	if ss.length == 0 {
		return nil
	}

	n := ss.top
	ss.top = n.prev
	ss.length--
	return n.value
}

// Push a value onto the top of the stack
func (ss *Stack) push(value interface{}) {
	n := &node{value, ss.top}
	ss.top = n
	ss.length++
}

/**

Implement a simple text editor. The editor initially contains an empty string, . Perform  operations of the following  types:

1. append - Append string  to the end of .
2. delete - Delete the last  characters of .
3. print - Print the  character of .
4. undo - Undo the last (not previously undone) operation of type  or , reverting  to the state it was in prior to that operation.

*/

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var stack = New()
	var result []string
	reader := bufio.NewReader(os.Stdin)
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	query := int(nTemp)
	i := 0
	for i < query {
		reader := bufio.NewReader(os.Stdin)
		str, _, _ := reader.ReadLine()
		split := strings.Split(string(str), " ")
		choice, _ := strconv.ParseInt(split[0], 10, 64)
		i++
		switch choice {
		case 1:
			if stack.isEmpty() {
				stack.push(split[1])
			} else {
				stack.push(fmt.Sprintf("%v%v", stack.peek(), split[1]))
			}
		case 2:
			if stack.isEmpty() {
				panic("stack is empty")
			}
			content := fmt.Sprint(stack.peek())
			k, _ := strconv.ParseInt(split[1], 10, 64)
			stack.push(fmt.Sprintf("%v", content[:len(content)-int(k)]))
		case 3:
			if stack.isEmpty() {
				panic("stack is empty")
			}
			content := fmt.Sprint(stack.peek())
			k, _ := strconv.ParseInt(split[1], 10, 64)
			fmt.Println(string(content[k-1]))
			result = append(result, string(content[k-1]))
		case 4:
			stack.pop()
		default:
			break
		}
	}

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
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

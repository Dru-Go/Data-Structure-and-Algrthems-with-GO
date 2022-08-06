package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	fmt.Println("s", s)
	fmt.Println("d", d)
	fmt.Println("m", m)

	var counter = 0
	for i := range s {
		if len(s) < i+int(m) {
			break
		}
		temp_list := s[i : int32(i)+m]
		reduced := reduce(temp_list, add, 0)
		fmt.Println("Reduced", reduced)
		if reduced == d {
			counter++
		}
	}

	return int32(counter)
}

func add(acc, curr int32) int32 {
	return curr + acc
}

func reduce(s []int32, f func(M, T int32) int32, initValue int32) int32 {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
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

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	dTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := birthday(s, d, m)

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

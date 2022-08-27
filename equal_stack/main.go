package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'equalStacks' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h1
 *  2. INTEGER_ARRAY h2
 *  3. INTEGER_ARRAY h3
 */

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	// Write your code here

	// var counter = 0
	if len(h1)+len(h2)+len(h3) == 0 {
		return 0
	}

	h1Stock, h2Stock, h3Stock := Sum(h1), Sum(h2), Sum(h3)

	for {
		min := MinIntSlice(int(h1Stock), int(h2Stock), int(h3Stock))

		if min == 0 {
			return 0
		}

		if min < int(h1Stock) {
			h1Stock -= h1[0]
		} else if min < int(h2Stock) {
			h2Stock -= h2[0]
		} else if min < int(h3Stock) {
			h3Stock -= h3[0]
		}

		if h1Stock == h2Stock && h2Stock == h3Stock {
			return h1Stock
		}
	}
}

func MinIntSlice(v ...int) int {
	sort.Ints(v)
	return v[0]
}

func Sum(array []int32) int32 {
	return reduce(array, sum, 0)
}

func sum(acc, current int32) int32 {
	return acc + current
}

func reduce(s []int32, f func(int32, int32) int32, initValue int32) int32 {
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

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n1 := int32(n1Temp)

	n2Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n2 := int32(n2Temp)

	n3Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n3 := int32(n3Temp)

	h1Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h1 []int32

	for i := 0; i < int(n1); i++ {
		h1ItemTemp, err := strconv.ParseInt(h1Temp[i], 10, 64)
		checkError(err)
		h1Item := int32(h1ItemTemp)
		h1 = append(h1, h1Item)
	}

	h2Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h2 []int32

	for i := 0; i < int(n2); i++ {
		h2ItemTemp, err := strconv.ParseInt(h2Temp[i], 10, 64)
		checkError(err)
		h2Item := int32(h2ItemTemp)
		h2 = append(h2, h2Item)
	}

	h3Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h3 []int32

	for i := 0; i < int(n3); i++ {
		h3ItemTemp, err := strconv.ParseInt(h3Temp[i], 10, 64)
		checkError(err)
		h3Item := int32(h3ItemTemp)
		h3 = append(h3, h3Item)
	}

	result := equalStacks(h1, h2, h3)

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

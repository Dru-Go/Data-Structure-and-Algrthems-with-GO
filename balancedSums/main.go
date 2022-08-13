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
 * Complete the 'balancedSums' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func balancedSums(arr []int32) string {
	if len(arr) == 1 {
		return "YES"
	}
	var left int32
	var right = reduce(arr, sum, 0)
	for i := 0; i < len(arr) && i+1 < len(arr); i++ {
		if i != 0 {
			left += arr[i-1]
		}
		right -= arr[i]
		fmt.Printf("%v = %v = %v \n", left, arr[i], right)
		if left == right {
			return "YES"
		}
	}
	return "NO"
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

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := balancedSums(arr)

		fmt.Fprintf(writer, "%s\n", result)
	}

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

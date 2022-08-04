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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	// count the positive, negative and zero elements
	var ratio map[int]int32 = map[int]int32{-1: 0, 0: 0, 1: 0}

	for _, v := range arr {
		switch {
		case len(arr) == 0:
			printer(ratio, len(arr))
			return
		case v < 0:
			ratio[-1]++
		case v == 0:
			ratio[0]++
		case v > 0:
			ratio[1]++
		}
	}
	printer(ratio, len(arr))
}

func printer(ratio map[int]int32, length int) {
	if length == 0 {
		fmt.Println(length)
		fmt.Println(length)
		fmt.Println(length)
	}
	fmt.Printf("%.6f \n", float32(ratio[1])/float32(length))
	fmt.Printf("%.6f \n", float32(ratio[-1])/float32(length))
	fmt.Printf("%.6f \n", float32(ratio[0])/float32(length))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

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

	plusMinus(arr)
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

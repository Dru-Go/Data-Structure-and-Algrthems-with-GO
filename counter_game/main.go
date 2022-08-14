package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'counterGame' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts LONG_INTEGER n as parameter.
 */

func counterGame(n int64) string {
	// Write your code here
	var Louise, Richard = "Louise", "Richard"
	var counter int64 = 1
	if n == 1 {
		return Richard
	}
	for n != 1 && n > 0 {
		if isPowerOfTwo(n) {
			n /= 2
		} else {
			n -= previousPowerOfTwo(n)
		}
		counter++
	}

	if counter%2 == 0 {
		return Louise
	} else {
		return Richard
	}
}

func previousPowerOfTwo(x int64) int64 {
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	return x - (x >> 1)
}

func isPowerOfTwo(n int64) bool {
	if n == 0 {
		return false
	}

	return (math.Ceil(math.Log2(float64(n))) == math.Floor(math.Log2(float64(n))))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := counterGame(n)

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

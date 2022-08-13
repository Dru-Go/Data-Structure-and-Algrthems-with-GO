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
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigit(n string, k int32) int32 {
	// Write your code here
	splits := strings.Split(n, "")
	sum := reduce(splits, k, sum, "")
	if len(sum) <= 1 {
		parsed, _ := strconv.Atoi(sum)
		return int32(parsed)
	}
	return superDigit(sum, 1)
}

func sum(acc, current string) string {
	ac, _ := strconv.Atoi(acc)
	cur, _ := strconv.Atoi(current)
	return fmt.Sprintf("%v", ac+cur)
}

func reduce(s []string, k int32, f func(string, string) string, initValue string) string {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	temp, _ := strconv.Atoi(acc)
	return fmt.Sprintf("%v", int64(temp)*int64(k))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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

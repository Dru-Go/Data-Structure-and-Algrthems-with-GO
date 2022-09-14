package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'noPrefix' function below.
 *
 * The function accepts STRING_ARRAY words as parameter.
 */

func noPrefix(words []string) string {
	// Write your code here
	return "BAD SET"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var words []string

	for i := 0; i < int(n); i++ {
		wordsItem := readLine(reader)
		words = append(words, wordsItem)
	}

	noPrefix(words)
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

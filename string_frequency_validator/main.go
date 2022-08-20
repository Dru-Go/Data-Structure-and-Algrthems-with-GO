package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	// Write your code here
	copy := s
	if len(s) == 0 {
		return "YES"
	}
	var counter []int
	for _, v := range copy {
		count := strings.Count(copy, string(v))
		if count != 0 {
			counter = append(counter, count)
		}
		copy = strings.ReplaceAll(copy, string(v), "")
	}

	if len(counter) == 1 {
		return "YES"
	}

	// sort counter
	sort.Slice(counter, func(i, j int) bool {
		return counter[i] < counter[j]
	})

	fmt.Println(counter)

	if counter[0] == counter[len(counter)-1] {
		return "YES"
	} else if counter[0] == 1 &&
		counter[1] == counter[len(counter)-1] {
		return "YES"
	} else if counter[0] == counter[1] &&
		counter[1] == counter[len(counter)-2] &&
		counter[len(counter)-2] == (counter[len(counter)-1]-1) {
		return "YES"
	} else {
		return "NO"
	}
}

func contains(s []int, searchNumbers int) bool {
	i := sort.SearchInts(s, searchNumbers)
	return i < len(s) && s[i] == searchNumbers
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

// Welcome
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

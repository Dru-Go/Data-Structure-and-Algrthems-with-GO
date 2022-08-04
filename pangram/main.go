package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	// Write your code here
	if checkPangram(s) {
		return "pangram"
	} else {
		return "not pangram"
	}
}

func checkPangram(s string) bool {
	set := make(map[int]string)
	var splattedStrings = strings.Split(s, "")

	for i := 0; i < len(splattedStrings); i++ {
		var str = []rune(strings.ToLower(splattedStrings[i]))
		var ascii = int(str[0])
		if (ascii == 32 || ascii >= 65 && ascii <= 90) || (ascii >= 97 && ascii <= 122) {
			set[ascii] = splattedStrings[i]
			continue
		} else {
			fmt.Printf("%v %v", str, ascii)
			return false
		}
	}
	fmt.Printf("%v", set)

	return len(set) == 27
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Fprintf(writer, "%s\n", result)

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

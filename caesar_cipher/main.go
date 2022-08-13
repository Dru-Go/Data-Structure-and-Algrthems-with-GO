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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

// Cipher formula
// c = (p + k) % 26 // p being the ascii of a byte in s

func caesarCipher(s string, k int32) string {
	// shift -> number of letters to move to right
	// offset -> size of the alphabet, in this case the plain ASCII
	shift, offset := rune(k), rune(26)

	runes := []rune(s)

	for index, char := range runes {
		if char >= 97 && char <= 122 {
			char = ((char-97+shift)%offset + 97)
		} else if char >= 65 && char <= 90 {
			char = ((char-65+shift)%offset + 65)
		}
		runes[index] = char
	}

	return string(runes)
}

func decipherCipher(s string, k int32) string {
	// shift -> number of letters to move to right
	// offset -> size of the alphabet, in this case the plain ASCII
	shift, offset := rune(k), rune(26)

	runes := []rune(s)

	for index, char := range runes {
		if char >= 97 && char <= 122 {
			r := (char - 97 - shift)
			if r < 0 {
				char = ((r)%offset + 123)
			} else {
				char = ((r)%offset + 97)
			}
		} else if char >= 65 && char <= 90 {
			r := (char - 65 - shift)
			if r < 0 {
				char = ((r)%offset + 91)
			} else {
				char = ((r)%offset + 65)
			}
		}
		runes[index] = char
	}

	return string(runes)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	// n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

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

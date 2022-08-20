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
 * Complete the 'bomberMan' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY grid
 */

func bomberMan(n int32, grid []string) []string {
	// Write your code here
	if n == 1 {
		return grid
	} else if n == 2 {
		var newGrid []string
		replacer := strings.NewReplacer(".", "0")
		for _, v := range grid {
			newGrid = append(newGrid, replacer.Replace(v))
		}
		return newGrid
	} else {
		var newGrid []string
		replacer := strings.NewReplacer(".", "0", "0", ".")
		for _, v := range grid {
			newGrid = append(newGrid, replacer.Replace(v))
		}
		for i, v := range newGrid {
			if strings.Contains(v, ".") {
				indexes := indexOfRune(v, byte('.'))
				for j := 0; j < len(indexes); j++ {
					if i-1 >= 0 {
						// top
						newGrid[i-1] = strings.Replace(newGrid[i-1], "0", ".", indexes[j])
					}
					if indexes[j]-1 >= 0 {
						// left
						newGrid[i] = strings.Replace(newGrid[i], "0", ".", indexes[j]-1)
					}
					if indexes[j]+1 <= len(grid[i]) {
						// right
						newGrid[i] = strings.Replace(newGrid[i], "0", ".", indexes[j]+1)
					}
					if i+1 <= len(newGrid)-1 {
						// bottom
						newGrid[i+1] = strings.Replace(newGrid[i+1], "0", ".", indexes[j])
					}
				}
			}

		}
		return newGrid
	}
}

func indexOfRune(str string, subString byte) []int {
	var indexes = []int{}
	for i := 0; i < len(str); i++ {
		if str[i] == subString {
			indexes = append(indexes, int(i))
		}
	}
	return indexes
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	_, err = strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	// c := int32(cTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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

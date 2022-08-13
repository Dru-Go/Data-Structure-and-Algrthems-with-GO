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
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

//  function gridChallenge(grid) {
// 	// create a sorted grid
// 	const sorted = grid.map((row) => row.split('').sort())
// 	// check that columns are sorted too
// 	for(let c = 0; c < grid.length; c++) {
// 		for(let r = 0; r < grid.length-1; r++) {
// 			if (sorted[r + 1][c] < sorted[r][c]) {
// 				return 'NO';
// 			}
// 		}
// 	}
// 	return 'YES'
// }

func gridChallenge(grid []string) string {
	var sorted []string
	if len(grid) <= 1 {
		return "NO"
	}
	for i := 0; i < len(grid); i++ {
		var current = strings.Split(grid[i], "")
		sort.Strings(current)
		sorted = append(sorted, strings.Join(current, ""))
	}

	for i := 0; i < len(grid) && i+1 < len(grid); i++ {
		first := sorted[i]
		second := sorted[i+1]
		fmt.Printf("sorted %v | %v \n", first, second)
		for j := 0; len(first) == len(second); j++ {
			if first[j] > second[j] {
				return "NO"
			}
		}
	}

	// Write your code here
	return "YES"
}

// sort.Strings(first)
// sort.Strings(second)
// fmt.Printf("sorted %v | %v \n", first, second)
// for j := 0; j < len(first) && j < len(second); j++ {
// 	if first[j] > second[j] {
// 		return "NO"
// 	}
// }

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
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var grid []string

		for i := 0; i < int(n); i++ {
			gridItem := readLine(reader)
			grid = append(grid, gridItem)
		}

		result := gridChallenge(grid)

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

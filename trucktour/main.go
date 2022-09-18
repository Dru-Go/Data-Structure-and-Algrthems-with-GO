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
 * Complete the 'truckTour' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY petrolpumps as parameter.
 */

func truckTour(petrolpumps [][]int32) int32 {
	// Write your code here
	var starting, diff int32 = 0, 0
	var cap int32 = 0
	for i := 0; i < len(petrolpumps); i++ {
		cap += petrolpumps[i][0] - petrolpumps[i][1]
		if cap < 0 {
			starting = int32(i + 1)
			diff += cap
			cap = 0
		}
	}
	if cap+diff >= 0 && len(petrolpumps) > 0 {
		return starting
	}
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var petrolpumps [][]int32
	for i := 0; i < int(n); i++ {
		petrolpumpsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var petrolpumpsRow []int32
		for _, petrolpumpsRowItem := range petrolpumpsRowTemp {
			petrolpumpsItemTemp, err := strconv.ParseInt(petrolpumpsRowItem, 10, 64)
			checkError(err)
			petrolpumpsItem := int32(petrolpumpsItemTemp)
			petrolpumpsRow = append(petrolpumpsRow, petrolpumpsItem)
		}

		if len(petrolpumpsRow) != 2 {
			panic("Bad input")
		}

		petrolpumps = append(petrolpumps, petrolpumpsRow)
	}

	result := truckTour(petrolpumps)

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

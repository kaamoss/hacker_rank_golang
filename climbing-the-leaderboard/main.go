package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func removeDuplicates(a []int) []int {
	r := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			r = append(r, val)
			seen[val] = val
		}
	}
	return r
}

func getMidPoint(start, end int) int {
	i := end - start
	if i%2 == 0 {
		return i/2 + start
	}
	return (i+1)/2 + start
}

func binarySearch(scores []int, v int, i int) int {
	var mid int
	var end int

	start := 0
	if i >= len(scores) {
		end = len(scores) - 1
	} else {
		end = i
	}

	for true {
		if end-start == 1 {
			// fmt.Printf("1 - v: %d start: %d end %d\n", v, start, end)
			if v < scores[end] {
				return end + 2
			} else if v < scores[start] {
				return start + 2
			}
			return 1
		}
		if end == 0 {
			return 1
		}

		mid = getMidPoint(start, end)

		if v > scores[mid] {
			end = mid
		} else if v < scores[mid] {
			start = mid
		} else {
			// v == scores[mid]
			// fmt.Printf("2 - v: %d start: %d end %d\n", v, start, end)
			return mid + 1
		}
	}
	return end
}

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int, alice []int) []int {
	r := make([]int, len(alice))

	//create scores and remove dups
	uniqueScores := removeDuplicates(scores)
	i := len(scores)

	for a, v := range alice {
		// fmt.Printf("1 - i: %d a: %d alice[a]: %d\n", i, a, alice[a])
		i = binarySearch(uniqueScores, v, i-1)
		r[a] = i
	}

	return r
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*3)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024*3)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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

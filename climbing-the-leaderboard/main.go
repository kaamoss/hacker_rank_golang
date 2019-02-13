package main

import (
"bufio"
"fmt"
"io"
"os"
"strconv"
"strings"
)

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	var aliceRanks = make([]int32, len(alice), len(alice))
	for foo, aliceScore := range alice {
		var rank int32
		var lastRank int32= 1
		var lastScore int32
		for _, leaderboardScore := range scores {
			if rank == 0 && lastScore > 0 && lastScore != leaderboardScore {
				lastRank += 1
			}

			if(rank == 0 && aliceScore >= leaderboardScore) {
				rank = lastRank
			}

			lastScore = leaderboardScore
		}

		if(rank == 0) {
			rank = lastRank+1
		}

		aliceRanks[foo] = rank
	}
	return aliceRanks
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024*3)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024*3)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
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



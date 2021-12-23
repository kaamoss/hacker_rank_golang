package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func arrayManipulation(n int32, queries [][]int32) int64 {
	//This is so slow, and why bother actually manipulating the array?
	var maxValue int64
	result := make([]int64, n)
	for _, query := range queries {
		a := query[0]
		b := query[1]
		k := int64(query[2])
		for i := a - 1; i < b; i++ {
			result[i] += k
			if result[i] > maxValue {
				maxValue = result[i]
			}
		}
	}

	//fmt.Println(result, n)
	return maxValue
}

func fasterArrayManip(n int32, queries [][]int32) int64 {
	var maxValue int64
	result := make([]int64, n)
	for _, query := range queries {
		a := query[0]
		b := query[1]
		k := int64(query[2])
		result[a-1] += k
		if b < n {
			result[b] -= k
		}
	}

	var runningTotal int64
	for _, val := range result {
		runningTotal += val
		if runningTotal > maxValue {
			maxValue = runningTotal
		}
	}

	return maxValue
}

func main() {
	file, err := os.Open("array-manipulation/insane_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	firstMultipleInput := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		scanner.Scan()
		queriesRowTemp := strings.Split(strings.TrimRight(scanner.Text(), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	//result := arrayManipulation(n, queries)
	result := fasterArrayManip(n, queries)

	fmt.Printf("%d\n", result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

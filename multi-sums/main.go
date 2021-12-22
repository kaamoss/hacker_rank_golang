package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.codechef.com/problems/FLOW001

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	totalSumsStr := scanner.Text()
	totalSums, err := strconv.Atoi(totalSumsStr)
	if err != nil {
		return
	}

	cntr := 0
	for scanner.Scan() {
		cntr++
		if cntr > totalSums {
			return
		}
		currentSum := 0
		numBatchStr := scanner.Text()
		termStrs := strings.Fields(numBatchStr)
		for _, termStr := range termStrs {
			termNum, err := strconv.Atoi(termStr)
			if err == nil {
				currentSum += termNum
			}
		}
		fmt.Println(currentSum)
	}
}

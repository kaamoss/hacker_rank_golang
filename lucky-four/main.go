package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.codechef.com/problems/LUCKFOUR

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numTotalIntsStr := scanner.Text()
	numTotalInts, err := strconv.Atoi(numTotalIntsStr)
	if err != nil {
		return
	}

	for i := 0; i < numTotalInts; i++ {
		scanner.Scan()
		currentNumStr := scanner.Text()
		fmt.Println(strings.Count(currentNumStr, "4"))
	}

}

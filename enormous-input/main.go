package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.codechef.com/problems/INTEST

func main() {
	var numDivisibleByK int = 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initialInputStr := scanner.Text()
	initialInputParts := strings.Split(initialInputStr, " ")
	if len(initialInputParts) != 2 {
		return
	}
	n, err := strconv.Atoi(initialInputParts[0])
	if err != nil {
		return
	}
	k, err := strconv.Atoi(initialInputParts[1])
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		currentLine := scanner.Text()
		currentNum, err := strconv.Atoi(currentLine)
		if err == nil && currentNum%k == 0 {
			numDivisibleByK++
		}
	}

	fmt.Println(numDivisibleByK)
}

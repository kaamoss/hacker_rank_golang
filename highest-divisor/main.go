package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//https://www.codechef.com/problems/HDIVISR

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputNumStr := scanner.Text()
	inputNum, err := strconv.Atoi(inputNumStr)
	if err != nil {
		return
	}
	fmt.Println(highestDivisor(inputNum))
}

func highestDivisor(num int) int {
	for i := 10; i > 0; i-- {
		if num%i == 0 {
			return i
		}
	}
	return 1
}

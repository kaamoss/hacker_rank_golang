package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.codechef.com/problems/HS08TEST

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	userInput := scanner.Text()
	userInputParts := strings.Split(userInput, " ")
	if len(userInputParts) != 2 {
		return
	}

	withdrawlAmt, err := strconv.Atoi(userInputParts[0])
	if err != nil {
		return
	}
	accountBalance, err := strconv.ParseFloat(userInputParts[1], 32)
	if err != nil {
		return
	}
	if withdrawlAmt%5 == 0 && (accountBalance-float64(withdrawlAmt)-0.5) >= 0 {
		accountBalance -= float64(withdrawlAmt) + .5
	}

	fmt.Printf("%.2f", accountBalance)
}

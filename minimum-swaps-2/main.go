package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	return doSwap(arr, 0)
}

func doSwap(arr []int32, numSwaps int32) int32 {
	for i, v := range arr {
		index := v - 1
		if v != arr[index] {
			arr[i], arr[index] = arr[index], arr[i]
			numSwaps++
			return doSwap(arr, numSwaps)
		}
	}
	fmt.Println(arr)
	return numSwaps
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nTemp, err := strconv.ParseInt(scanner.Text(), 10, 64)
	checkError(err)
	n := int32(nTemp)

	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		valueItemTemp, err := strconv.ParseInt(values[i], 10, 64)
		checkError(err)
		valueItem := int32(valueItemTemp)
		arr = append(arr, valueItem)
	}

	res := minimumSwaps(arr)

	fmt.Printf("%d\n", res)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

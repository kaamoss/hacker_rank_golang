package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

//https://www.codechef.com/problems/TSORT

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initialInput := scanner.Text()
	numItems, err := strconv.Atoi(initialInput)
	if err != nil {
		return
	}

	allItems := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		scanner.Scan()
		currentItemStr := scanner.Text()
		currentItem, err := strconv.Atoi(currentItemStr)
		if err == nil {
			allItems[i] = currentItem
		}
	}
	sort.Ints(allItems)

	for _, item := range allItems {
		fmt.Println(item)
	}
}

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
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
	totalBribes := 0
	for i := len(q) - 1; i > 0; i-- {
		if q[i] != int32(i+1) {
			numBribed := 0
			if i-1 >= 0 && q[i-1] == int32(i+1) {
				numBribed = 1
				q[i], q[i-1] = q[i-1], q[i]
			} else if i-2 >= 0 && q[i-2] == int32(i+1) {
				numBribed = 2
				q[i], q[i-1], q[i-2] = q[i-2], q[i], q[i-1]
			} else if i+1 < len(q) && q[i+1] == int32(i+1) {
				numBribed = 1
				q[i+1], q[i] = q[i], q[i+1]
			} else if i+2 < len(q) && q[i+2] == int32(i+1) {
				numBribed = 2
				q[i+2], q[i+1], q[i] = q[i+1], q[i], q[i+2]
			} else {
				fmt.Println("Too chaotic")
				return
			}
			totalBribes += numBribed
		}
	}
	fmt.Println(totalBribes)
}

func main() {
	//expected output of input1.txt: 966
	/* expected output of input2.txt:

	1201
	1208
	Too chaotic
	1196
	Too chaotic
	1196
	Too chaotic
	Too chaotic
	Too chaotic
	1210
	*/
	file, err := os.Open("new-years-chaos/input2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	tTemp, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		scanner.Scan()
		nTemp, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
		checkError(err)
		n := int32(nTemp)

		scanner.Scan()
		qTemp := strings.Fields(scanner.Text())

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

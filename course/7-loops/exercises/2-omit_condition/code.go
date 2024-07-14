package main

import (
	"fmt"
)

const baseCost = 1.0

func maxMessages(thresh float64) int {
	var totalCost float64
	for i := 0; ; i++ {
		totalCost += (0.01 * float64(i)) + baseCost
		if totalCost > thresh {
			return i
		}
	}
}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(40.00)
	test(50.00)
}

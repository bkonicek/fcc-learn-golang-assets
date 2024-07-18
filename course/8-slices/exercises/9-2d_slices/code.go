package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	gridRow := make([][]int, rows)
	for x := 0; x < rows; x++ {
		gridRow[x] = make([]int, cols)
		for y := 0; y < cols; y++ {
			gridRow[x][y] = x * y
		}
	}

	return gridRow
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}

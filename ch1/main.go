package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findSmallestNumber([]int{2, 4, 1}))

	fmt.Println(findGreatesCommonDivisor(54, 24))

	fmt.Println(findSquareRoot(25.25))

	fmt.Println(findElementInArray([]int{1, 5, 24, 11, -1}, 11))

	fmt.Println(findMaxInArray([]int{2, 4, 1}))

	fmt.Println(findMinInArray([]int{2, 4, 1}))

	fmt.Println(find2ndMaxInArray([]int{2, 4, 1, 10, 5}))

	fmt.Println(sortArray([]int{2, 4, 1, 10, 5}))

	fmt.Println(findMaxInMatrix([][]int{
		[]int{10, 20, 30},
		[]int{11, 50, 8},
		[]int{14, 100, 43},
	}))

	fmt.Println(findTraceInMatrix([][]int{
		[]int{10, 20, 30},
		[]int{11, 50, 8},
		[]int{14, 100, 43},
	}))
}

func findSmallestNumber(input []int) int {
	sort.Ints(input)

	return input[0]
}

func findGreatesCommonDivisor(input1, input2 int) (result []int) {
	var max int
	if max = input1; max < input2 {
		max = input2
	}

	for i := 1; i < max; i++ {
		if input1%i == 0 && input2%i == 0 {
			result = append(result, i)
		}
	}

	return
}

func findSquareRoot(input float64) (result float64) {
	return math.Sqrt(input)
}

func findElementInArray(input []int, find int) bool {
	for _, number := range input {
		if find == number {
			return true
		}
	}

	return false
}

func findMaxInArray(input []int) (result int) {
	sort.Ints(input)

	return input[len(input)-1]
}

func findMinInArray(input []int) (result int) {
	sort.Ints(input)

	return input[0]
}

func find2ndMaxInArray(input []int) (result int) {
	sort.Ints(input)

	return input[len(input)-2]
}

func sortArray(input []int) []int {
	sort.Ints(input)

	return input
}

func findMaxInMatrix(input [][]int) (result int) {
	for i, row := range input {
		for j := range row {
			if (i == 0 && j == 0) || result < input[i][j] {
				result = input[i][j]
			}
		}
	}

	return
}

func findTraceInMatrix(input [][]int) (result int) {
	for i, row := range input {
		for j := range row {
			if i == j {
				result += input[i][j]
			}
		}
	}

	return
}

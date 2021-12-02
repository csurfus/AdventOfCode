package main

import (
	"fmt"
)

func starOne(input []int) {
	var sum = 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i - 1] {
			sum += 1
		}
	}
	fmt.Println(sum, "measurements are larger than the previous measurement")
}

func starTwo(input []int) {
	var sum = 0
	var previousSum = input[0] + input[1] + input[2]
	var currentSum = 0
	for i := 1; i < len(input) - 2; i++ {
		currentSum = input[i] + input[i + 1] + input[i + 2]
		if currentSum > previousSum {
			sum += 1
		}
		previousSum = currentSum
	}
	fmt.Println(sum, "sums are larger than the previous sum")
}

func main() {
	input := adventOfCodeInputReader("./inputs/day1.txt")

	starOne(input)
	starTwo(input)
}

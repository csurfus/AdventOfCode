package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func adventOfCodeInput(path string) []int {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("ERROR READING FILE", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []int
	scanner.Scan()

	var inputNumbers = strings.Split(scanner.Text(), ",")

	for i := 0; i < len(inputNumbers); i++ {
		entry, _ := strconv.Atoi(inputNumbers[i])
		entries = append(entries, entry)
	}

	return entries
}

func MinInt(one int, two int) int {
	if one < two {
		return one
	}

	return two
}

func AbsInt(input int) int {
	if input < 0 {
		return -input
	}

	return input
}

func day7starOne(input []int) {
	var sum = math.MaxInt
	for i := 0; i < len(input); i++ {
		localSum := 0
		for j := 0; j < len(input); j++ {
			if j == i {
				continue
			}

			localSum += AbsInt(input[i] - input[j])
		}
		sum = MinInt(localSum, sum)
	}

	fmt.Println("Minimum Amount of Fuel", sum)
}

func countSteps(input int) int {
	var sum = 0
	for i := 1; i <= input; i++ {
		sum += i
	}
	return sum
}

func average(input []int) float64 {
	var sum = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	temp := float64(sum) / float64(len(input))
	return temp
}

func calculateSum(input []int, average float64) int {
	var sum = 0
	for i := 0; i < len(input); i++ {
		temp := AbsInt(int(average) - input[i])
		steps := countSteps(temp)
		sum += steps
	}
	return sum
}

func day7starTwo(input []int) {
	var average = average(input)
	var averageFloor = calculateSum(input, math.Floor(average))
	var averageCeil = calculateSum(input, math.Ceil(average))
	sum := MinInt(averageFloor, averageCeil)

	fmt.Println("Minimum Amount of Fuel", sum)
}

func main() {
	input := adventOfCodeInput("./inputs/day7.txt")
	day7starOne(input)
	day7starTwo(input)
}
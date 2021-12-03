package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func adventOfCodeInput(path string) []string {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("ERROR READING FILE", err)
		return nil
	}

	defer file.Close()

	var binaryNumbers []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		binaryNumbers = append(binaryNumbers, scanner.Text())
	}

	return binaryNumbers
}


func day3starOne(binaryNumbers []string) {
	var common []string
	var least []string

	for i:= 0; i < len(binaryNumbers[0]); i++ {
		trueCount := 0
		falseCount := 0
		for j:= 0; j < len(binaryNumbers); j++ {
			if string(binaryNumbers[j][i]) == "0" {
				falseCount++
			}

			if string(binaryNumbers[j][i]) == "1" {
				trueCount++
			}
		}
		if trueCount > falseCount {
			common = append(common, "1")
			least = append(least, "0")
		} else {
			common = append(common, "0")
			least = append(least, "1")
		}
	}

	gammaRateDecimal, _ := strconv.ParseInt(strings.Join(common, ""), 2, 64)
	epsilonRateDecimal, _ := strconv.ParseInt(strings.Join(least, ""), 2, 64)
	powerConsumption  := gammaRateDecimal * epsilonRateDecimal

	fmt.Println("Gamma Rate:", common, "Decimal:", gammaRateDecimal)
	fmt.Println("Epsilon Rate:", least, "Decimal:", epsilonRateDecimal)
	fmt.Println("Power Consumption:", powerConsumption)
}

func removeElementsWith(criteria string, position int, input []string) []string {
	if len(input) == 1 { return input }

	var inputCopy []string
	for i := 0; i < len(input); i++ {
		if string(input[i][position]) != criteria {
			inputCopy = append(inputCopy, input[i])
		}
	}
	return inputCopy
}

func algo(input []string, priority string, notPriority string) []string {
	for i:= 0; i < len(input[0]); i++ {
		trueCount := 0
		falseCount := 0
		for j:= 0; j < len(input); j++ {
			if string(input[j][i]) == "0" {
				falseCount++
			}

			if string(input[j][i]) == "1" {
				trueCount++
			}
		}
		if trueCount >= falseCount {
			input = removeElementsWith(priority, i, input)
		} else {
			input = removeElementsWith(notPriority, i, input)
		}
	}

	return input
}

func day3starTwo(binaryNumbers []string) {
	oxygenGeneratorDecimal, _ := strconv.ParseInt(strings.Join(algo(binaryNumbers, "1", "0"), ""), 2, 64)
	c02ScrubberDecimal, _ := strconv.ParseInt(strings.Join(algo(binaryNumbers, "0", "1"), ""), 2, 64)
	lifeSupportRating  := oxygenGeneratorDecimal * c02ScrubberDecimal

	fmt.Println("Oxygen Generator Rating", oxygenGeneratorDecimal)
	fmt.Println("C02 Scrubber Rating", c02ScrubberDecimal)
	fmt.Println("Life Support Rating", lifeSupportRating)
}

func main() {
	binaryNumbers := adventOfCodeInput("./inputs/day3.txt")

	day3starOne(binaryNumbers)
	day3starTwo(binaryNumbers)
}

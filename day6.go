package main

import (
	"bufio"
	"fmt"
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

func countFish(input []int) int {
	return len(input)
}

func countFishLarge(dayList map[int]int) int {
	var sum = 0

	for _, value := range dayList {
		sum += value
	}

	return sum
}

func day6starOne(input []int, days int) {
	temp := make([]int, len(input))
	copy(temp, input)

	for day := 1; day <= days; day++ {
		size := len(temp)
		for j := 0; j < size; j++ {
			if temp[j] == 0 {
				temp[j] = 6
				temp = append(temp, 8)
			} else {
				temp[j]--
			}
		}
		//fmt.Println("Day", day, input)
	}

	fmt.Println("Number of Fish", countFish(temp))
}

func day6starTwo(input []int, days int) {
	//Using days into buckets method
	dayList := make(map[int]int)
	for i := 0; i < len(input); i++ {
		num := input[i]
		dayList[num]++
	}

	for day := 0; day < days; day++ {
		cache := dayList[0]
		for i := 0; i < 8; i++ {
			dayList[i] = dayList[i + 1]
		}
		dayList[6] += cache
		dayList[8] = cache
	}

	fmt.Println("Number of Fish", countFishLarge(dayList))
}

func main() {
	input := adventOfCodeInput("./inputs/day6.txt")
	day6starOne(input, 80)
	day6starTwo(input, 256)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func adventOfCodeInput(path string) ([]string, []int) {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("ERROR READING FILE", err)
		return nil, nil
	}

	defer file.Close()

	var directions []string
	var nums []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitTemp := strings.Split(scanner.Text(), " ")

		direction := splitTemp[0]
		num, _ := strconv.Atoi(splitTemp[1])

		directions = append(directions, direction)
		nums = append(nums, num)
	}

	return directions, nums
}


func day2starOne(directions []string, nums []int) {
	var horizontalPos = 0
	var depth = 0

	for i := 0; i < len(directions); i++ {
		if directions[i] == "forward" {
			horizontalPos += nums[i]
		}

		if directions[i] == "down" {
			depth += nums[i]
		}

		if directions[i] == "up" {
			depth -= nums[i]
		}
	}
	fmt.Println("Horizontal Position:", horizontalPos)
	fmt.Println("Depth:", depth)
	fmt.Println("Multiplied:", horizontalPos * depth)
}

func day2starTwo(directions []string, nums []int) {
	var horizontalPos = 0
	var depth = 0
	var aim = 0

	for i := 0; i < len(directions); i++ {
		if directions[i] == "forward" {
			horizontalPos += nums[i]
			depth += aim * nums[i]
		}

		if directions[i] == "down" {
			aim += nums[i]
		}

		if directions[i] == "up" {
			aim -= nums[i]
		}
	}
	fmt.Println("Horizontal Position:", horizontalPos)
	fmt.Println("Depth:", depth)
	fmt.Println("Multiplied:", horizontalPos * depth)
}

func main() {
	directions, nums := adventOfCodeInput("./inputs/day2.txt")

	day2starOne(directions, nums)
	day2starTwo(directions, nums)
}

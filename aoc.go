package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func adventOfCodeInputReader(path string) []int {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("ERROR READING FILE", err)
		return nil
	}

	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, temp)
	}

	return lines
}

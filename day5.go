package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end point
}

func adventOfCodeInput(path string) []line {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("ERROR READING FILE", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []line

	for scanner.Scan() {
		var input = strings.ReplaceAll(scanner.Text(), " ", "")
		var bingoNumbers = strings.Split(input, "->")

		var firstPair = strings.Split(bingoNumbers[0], ",")
		firstX, _ := strconv.Atoi(firstPair[0])
		firstY, _ := strconv.Atoi(firstPair[1])

		var secondPair = strings.Split(bingoNumbers[1], ",")
		secondX, _ := strconv.Atoi(secondPair[0])
		secondY, _ := strconv.Atoi(secondPair[1])

		entry := line{
			start: point{
				x: firstX,
				y: firstY,
			},
			end: point{
				x: secondX,
				y: secondY,
			},
		}
		entries = append(entries, entry)
	}

	return entries
}

func getMaxes(lines []line) (int, int) {
	maxRow := 0
	maxCol := 0
	//col, row
	for i := 0; i < len(lines); i++ {
		maxRow = int(math.Max(float64(lines[i].start.y), float64(maxRow)))
		maxRow = int(math.Max(float64(lines[i].end.y), float64(maxRow)))
		maxCol = int(math.Max(float64(lines[i].start.x), float64(maxCol)))
		maxCol = int(math.Max(float64(lines[i].end.x), float64(maxCol)))
	}
	maxRow++
	maxCol++
	fmt.Println("Max Col", maxCol, "Max Row", maxRow)
	return maxCol, maxRow
}

func createBoard(maxCol int, maxRow int) [][]string {
	board := make([][]string, maxCol)
	for i := 0; i < maxCol; i++ {
		board[i] = make([]string, maxRow)
	}

	for i := 0; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {
			board[i][j] = "."
		}
	}

	return board
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}
}

func drawLine(board [][]string, entry line) [][]string {
	if entry.start.x != entry.end.x && entry.start.y != entry.end.y {
		return board
	}

	x := entry.start.x
	y := entry.start.y
	//fmt.Println("Start", entry.start)
	//fmt.Println("End", entry.end)
	for true {
		if board[y][x] != "." {
			temp, _ := strconv.Atoi(board[y][x])
			temp++
			board[y][x] = strconv.Itoa(temp)
		} else {
			board[y][x] = "1"
		}

		if x == entry.end.x && y == entry.end.y {
			return board
		}

		if entry.start.x < entry.end.x {
			x++
		}

		if entry.start.x > entry.end.x {
			x--
		}

		if entry.start.y < entry.end.y {
			y++
		}

		if entry.start.y > entry.end.y {
			y--
		}
	}
	return nil
}

func drawLineWithDiagonals(board [][]string, entry line) [][]string {
	x := entry.start.x
	y := entry.start.y
	//fmt.Println("Start", entry.start)
	//fmt.Println("End", entry.end)
	for true {
		if board[y][x] != "." {
			temp, _ := strconv.Atoi(board[y][x])
			temp++
			board[y][x] = strconv.Itoa(temp)
		} else {
			board[y][x] = "1"
		}

		if x == entry.end.x && y == entry.end.y {
			return board
		}

		if entry.start.x < entry.end.x {
			x++
		}

		if entry.start.x > entry.end.x {
			x--
		}

		if entry.start.y < entry.end.y {
			y++
		}

		if entry.start.y > entry.end.y {
			y--
		}
	}
	return nil
}

func countOverlaps(board [][]string) {
	overlaps := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != "." &&  board[i][j] != "1" {
				overlaps++
			}
		}
	}

	fmt.Println("Overlaps", overlaps)
}

func day5starOne(lines []line) {
	maxCol, maxRow := getMaxes(lines)
	board := createBoard(maxCol, maxRow)

	for i := 0; i < len(lines); i++ {
		board = drawLine(board, lines[i])
	}
	//printBoard(board)
	countOverlaps(board)
}

func day5starTwo(lines []line) {
	maxCol, maxRow := getMaxes(lines)
	board := createBoard(maxCol, maxRow)

	for i := 0; i < len(lines); i++ {
		board = drawLineWithDiagonals(board, lines[i])
	}
	//printBoard(board)
	countOverlaps(board)
}

func main() {
	lines := adventOfCodeInput("./inputs/day5.txt")

	day5starOne(lines)
	day5starTwo(lines)
}

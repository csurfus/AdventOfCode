/*
 * No clean code here
 * YIKES! Probably the most ugly code ever.
 * Starting to regret doing all of these in Go
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func adventOfCodeInput(path string) ([]string, [][][]string) {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("ERROR READING FILE", err)
		return nil, nil
	}

	defer file.Close()

	splitFn := func(c rune) bool {
		return c == ' '
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var bingoNumbers = strings.Split(scanner.Text(), ",")
	var boards [][][]string
	for scanner.Scan() {
		var board [][]string
		scanner.Scan()
		board = append(board, strings.FieldsFunc(scanner.Text(), splitFn))
		scanner.Scan()
		board = append(board, strings.FieldsFunc(scanner.Text(), splitFn))
		scanner.Scan()
		board = append(board, strings.FieldsFunc(scanner.Text(), splitFn))
		scanner.Scan()
		board = append(board, strings.FieldsFunc(scanner.Text(), splitFn))
		scanner.Scan()
		board = append(board, strings.FieldsFunc(scanner.Text(), splitFn))

		boards = append(boards, board)
	}

	return bingoNumbers, boards
}

func countUnmarkedNumbers(board [][]string) int {
	sum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != "X" {
				temp, _ := strconv.Atoi(board[i][j])
				sum += temp
			}
		}
	}

	return sum
}

func checkBoardsForWinner(drawnNumber string, board [][]string) (bool, [][]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if drawnNumber == board[i][j] {
				board[i][j] = "X"
				if isWinner(board) {
					return true, board
				}
			}
		}
	}
	return false, nil
}

func isWinner(board [][]string) bool {
	var numberOfX = 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			check := board[j][i]
			if check == "X" {
				numberOfX++
				if numberOfX == len(board) {
					return true
				}
			}
		}
		numberOfX = 0
	}

	numberOfX = 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			check := board[i][j]
			if check == "X" {
				numberOfX++
				if numberOfX == len(board) {
					return true
				}
			}
		}
		numberOfX = 0
	}

	return false
}

func day4starOne(bingoNumbers []string, boards [][][]string) {
	for i := 0; i < len(bingoNumbers); i++ {
		drawnNumber := bingoNumbers[i]
		for j := 0; j < len(boards); j++ {
			var isWin, winningBoard = checkBoardsForWinner(drawnNumber, boards[j])

			if isWin {
				fmt.Println("WINNER!")
				fmt.Println(winningBoard)
				sumOfAll := countUnmarkedNumbers(winningBoard)
				fmt.Println("Sum of all unmarked numbers:",  sumOfAll)
				fmt.Println("Number Drawn:",  drawnNumber)
				drawn, _ := strconv.Atoi(drawnNumber)
				finalScore := drawn * sumOfAll
				fmt.Println("Final Score", finalScore)

				return
			}
		}
	}
}

func day4starTwo(bingoNumbers []string, boards [][][]string) {
	for i := 0; i < len(bingoNumbers); i++ {
		drawnNumber := bingoNumbers[i] //Call number
		for j := 0; j < len(boards); j++ {
			var isWin, winningBoard = checkBoardsForWinner(drawnNumber, boards[j])

			if isWin {
				sumOfAll := countUnmarkedNumbers(winningBoard)
				fmt.Println("Sum of all unmarked numbers:",  sumOfAll)
				fmt.Println("Number Drawn:",  drawnNumber)
				drawn, _ := strconv.Atoi(drawnNumber)
				finalScore := drawn * sumOfAll
				fmt.Println("Final Score", finalScore)

				temp := [][]string {
					{ "-1", "-1", "-1", "-1", "-1" },
					{ "-1", "-1", "-1", "-1", "-1" },
					{ "-1", "-1", "-1", "-1", "-1" },
					{ "-1", "-1", "-1", "-1", "-1" },
					{ "-1", "-1", "-1", "-1", "-1" },
				}
				boards[j] = temp

			}
		}
	}
}

func main() {
	bingoNumbers, boards := adventOfCodeInput("./inputs/day4.txt")

	day4starOne(bingoNumbers, boards)
	day4starTwo(bingoNumbers, boards)
}

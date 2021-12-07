package helpers

import "fmt"

func Print2DInput(input [][]string) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			fmt.Print(input[i][j])
		}
		fmt.Println()
	}
}
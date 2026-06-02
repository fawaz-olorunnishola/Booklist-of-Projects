package main

import (
	"fmt"
	"os"
)

func parseArgs(grid *[9][9]byte) int {
	if len(os.Args) != 10 {
		return 0
	}

	for row := 0; row < 9; row++ {
		line := os.Args[row+1]

		if len(line) != 9 {
			return 0
		}

		for col := 0; col < 9; col++ {
			cell := line[col]

			if cell != '.' && (cell < '1' || cell > '9') {
				return 0
			}

			grid[row][col] = cell
		}
	}

	return 1
}

func printGrid(grid *[9][9]byte) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%c", grid[row][col])

			if col < 8 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func initialValid(grid *[9][9]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			value := grid[row][col]

			if value != '.' {

				grid[row][col] = '.'

				if !isValid(grid, row, col, value) {
					grid[row][col] = value
					return false
				}

				grid[row][col] = value
			}
		}
	}

	return true
}

func main() {
	var grid [9][9]byte

	if parseArgs(&grid) == 0 {
		fmt.Println("Error")
		return
	}

	if !initialValid(&grid) {
		fmt.Println("Error")
		return
	}

	if solve(&grid) == 0 {
		fmt.Println("Error")
		return
	}

	printGrid(&grid)
}
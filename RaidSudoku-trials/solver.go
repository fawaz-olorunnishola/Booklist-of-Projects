package main

func isValid(grid *[9][9]byte, row int, col int, value byte) bool {

	for currentCol := 0; currentCol < 9; currentCol++ {
		if grid[row][currentCol] == value {
			return false
		}
	}

	for currentRow := 0; currentRow < 9; currentRow++ {
		if grid[currentRow][col] == value {
			return false
		}
	}

	boxStartRow := row - row%3
	boxStartCol := col - col%3

	for rowOffset := 0; rowOffset < 3; rowOffset++ {
		for colOffset := 0; colOffset < 3; colOffset++ {

			if grid[boxStartRow+rowOffset][boxStartCol+colOffset] == value {
				return false
			}
		}
	}

	return true
}

func solve(grid *[9][9]byte) int {

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			if grid[row][col] == '.' {

				for value := byte('1'); value <= '9'; value++ {

					if isValid(grid, row, col, value) {

						grid[row][col] = value

						if solve(grid) == 1 {
							return 1
						}

						grid[row][col] = '.'
					}
				}

				return 0
			}
		}
	}

	return 1
}
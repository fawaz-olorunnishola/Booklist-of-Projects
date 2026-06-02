package main

import (
	"bufio"
	"fmt"
	"os"
)

func quadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == y && col == 1) {
				result += "o"
			} else if (row == 1 && col == x) || (row == y && col == x) {
				result += "o"
			} else if col == 1 || col == x {
				result += "|"
			} else if row == 1 || row == y {
				result += "-"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "/"
			} else if row == 1 && col == x {
				result += "\\"
			} else if row == y && col == 1 {
				result += "\\"
			} else if row == y && col == x {
				result += "/"
			} else if col == 1 || col == x || row == 1 || row == y {
				result += "*"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "A"
			} else if row == 1 && col == x {
				result += "A"
			} else if row == y && col == 1 {
				result += "C"
			} else if row == y && col == x {
				result += "C"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "A"
			} else if row == y && col == 1 {
				result += "A"
			} else if row == 1 && col == x {
				result += "C"
			} else if row == y && col == x {
				result += "C"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "A"
			} else if row == y && col == 1 {
				result += "C"
			} else if row == 1 && col == x {
				result += "C"
			} else if row == y && col == x {
				result += "A"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	y := len(lines)
	x := len([]rune(lines[0]))

	input := ""
	for i, line := range lines {
		input += line
		if i < len(lines)-1 {
			input += "\n"
		}
	}
	input += "\n"

	var matches []string

	if quadA(x, y) == input {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if quadB(x, y) == input {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if quadC(x, y) == input {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if quadD(x, y) == input {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if quadE(x, y) == input {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		result := ""
		for i, match := range matches {
			result += match
			if i < len(matches)-1 {
				result += " || "
			}
		}
		fmt.Println(result)
	}
}

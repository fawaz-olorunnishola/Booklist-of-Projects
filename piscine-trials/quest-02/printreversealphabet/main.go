package main

import "github.com/01-edu/z01"

func main() {
	for name := 'z'; name >= 'a'; name-- {
		z01.PrintRune(name)
	}
	z01.PrintRune('\n')
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func Sort(s string) string {
	r := []rune(s)

	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}

	return string(r)
}

func Help() {
	Print("--insert\n")
	Print("  -i\n")
	Print("\t This flag inserts the string into the string passed as argument.\n")
	Print("--order\n")
	Print("  -o\n")
	Print("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		Help()
		return
	}

	insert := ""
	order := false
	str := ""

	for i := 0; i < len(args); i++ {
		if args[i] == "--order" || args[i] == "-o" {
			order = true
		} else if args[i] == "--insert" || args[i] == "-i" {
			if i+1 < len(args) {
				insert = args[i+1]
				i++
			}
		} else if len(args[i]) > 9 && args[i][:9] == "--insert=" {
			insert = args[i][9:]
		} else if len(args[i]) > 3 && args[i][:3] == "-i=" {
			insert = args[i][3:]
		} else {
			str = args[i]
		}
	}

	result := str + insert

	if order {
		result = Sort(result)
	}

	Print(result)
	z01.PrintRune('\n')
}

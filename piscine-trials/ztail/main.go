package main

import (
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		os.Exit(1)
	}

	n := atoi(os.Args[2])
	files := os.Args[3:]
	hasErr := false

	for i, name := range files {
		data, err := os.ReadFile(name)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			hasErr = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				os.Stdout.WriteString("\n")
			}
			os.Stdout.WriteString("==> " + name + " <==\n")
		}

		if n > len(data) {
			n = len(data)
		}

		os.Stdout.Write(data[len(data)-n:])
	}

	if hasErr {
		os.Exit(1)
	}
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}

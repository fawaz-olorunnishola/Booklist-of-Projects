package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}

		io.Copy(os.Stdout, file)
		file.Close()
	}
}

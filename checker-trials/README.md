# Quadchecker
## Description

quadchecker is a Go program that reads an ASCII rectangle from standard input and determines which quad function (quadA–quadE) generated it.

It outputs the matching quad name(s) and their dimensions. If no match is found, it prints:

Not a quad function
Usage

## Run the program:

go run .
Example
./quadA 3 3 | go run .

Output:

[quadA] [3] [3]

## Multiple Matches
./quadC 1 1 | go run .

## Output:

[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]

## How It Works
- Reads input from stdin
- Determines width (x) and height (y)
- Compares input with all quad functions
- Prints matches in alphabetical order

## Structure
quadchecker/
├── go.mod
└── main.go

## Notes
- Output always ends with a newline
- Multiple matches are separated by ||
- Invalid input → Not a quad function 

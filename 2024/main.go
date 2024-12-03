package main

import (
	"aoc/calendar"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: go run aoc.go [stage] [input_file]")
		return
	}

	stage := args[1]
	file_name := args[2]

	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	switch stage {
	case "1a":
		fallthrough
	case "01a":
		calendar.Day01_Part1(lines)
	case "1b":
		fallthrough
	case "01b":
		calendar.Day01_Part2(lines)
	default:
		fmt.Printf("Stage `%s' not implemented\n", stage)
	}
}

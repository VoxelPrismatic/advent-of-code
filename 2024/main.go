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
	for lines[len(lines)-1] == "" {
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

	case "2a":
		fallthrough
	case "02a":
		calendar.Day02_Part1(lines)
	case "2b":
		fallthrough
	case "02b":
		calendar.Day02_Part2(lines)

	case "3a":
		fallthrough
	case "03a":
		calendar.Day03_Part1(lines)
	case "3b":
		fallthrough
	case "03b":
		calendar.Day03_Part2(lines)

	case "4a":
		fallthrough
	case "04a":
		calendar.Day04_Part1(lines)
	case "4b":
		fallthrough
	case "04b":
		calendar.Day04_Part2(lines)

	case "5a":
		fallthrough
	case "05a":
		calendar.Day05_Part1(lines)
	case "5b":
		fallthrough
	case "05b":
		calendar.Day05_Part2(lines)
	default:
		fmt.Printf("Stage `%s' not implemented\n", stage)
	}
}

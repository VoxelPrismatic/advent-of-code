package main

import (
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
	fmt.Println(lines)

	switch stage {
	default:
		fmt.Printf("Stage `%s' not implemented\n", stage)
	}
}

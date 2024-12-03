package calendar

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type day01_pair struct {
	left  int
	right int
}

func day01_pairs(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		ints := []int{}
		for _, part := range parts {
			if part == "" {
				continue
			}

			i, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			ints = append(ints, i)
		}

		if len(ints) != 2 {
			panic("invalid input")
		}

		left = append(left, ints[0])
		right = append(right, ints[1])
	}

	return left, right
}

func Day01_Part1(lines []string) {
	left, right := day01_pairs(lines)

	slices.Sort(left)
	slices.Sort(right)

	pairs := []day01_pair{}
	for i := 0; i < len(left); i++ {
		pairs = append(pairs, day01_pair{left[i], right[i]})
	}

	ret := 0
	for _, pair := range pairs {
		if pair.left >= pair.right {
			ret += pair.left - pair.right
		} else {
			ret += pair.right - pair.left
		}
	}

	fmt.Println(ret)
}

func Day01_Part2(lines []string) {
	left, right := day01_pairs(lines)
	counts := map[int]int{}

	ret := 0
	for _, i := range left {
		count, ok := counts[i]
		if !ok {
			count = 0
			for _, j := range right {
				if i == j {
					count++
				}
			}

			counts[i] = count
		}

		ret += i * count
	}

	fmt.Println(ret)
}

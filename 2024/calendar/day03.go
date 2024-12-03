package calendar

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day03_Part1(lines []string) {
	mul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	ret := 0

	for _, line := range lines {
		for _, match := range mul.FindAllStringSubmatch(line, -1) {
			left, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}

			right, err := strconv.Atoi(match[2])
			if err != nil {
				continue
			}

			ret += left * right
		}
	}

	fmt.Println(ret)
}

func Day03_Part2(lines []string) {
	line := strings.Join(lines, "")
	dont := regexp.MustCompile(`don't\(\).*?do\(\)`)
	dont_end := regexp.MustCompile(`don't\(\).*?$`)
	line = dont.ReplaceAllString(line, "")
	line = dont_end.ReplaceAllString(line, "")

	Day03_Part1([]string{line})
}

package calendar

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day05_GetRules(lines []string) map[int][]int {
	rules := map[int][]int{}
	for _, line := range lines {
		if !strings.Contains(line, "|") {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			panic("invalid input")
		}

		parent, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(err)
		}

		child, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic(err)
		}

		rules[parent] = append(rules[parent], child)
	}

	return rules
}

func Day05_ValidRule(rules map[int][]int, order []int) (bool, int) {
	rule, ok := rules[order[0]]
	if !ok {
		panic("cannot find rule for page " + strconv.Itoa(order[0]))
	}
	for _, page := range order[1:] {
		if !slices.Contains(rule, page) {
			return false, page
		}

		rule = rules[page]
	}

	return true, -1
}

func Day05_GetOrders(lines []string) [][]int {
	ordering := [][]int{}
	for _, line := range lines {
		if !strings.Contains(line, ",") {
			continue
		}

		parts := strings.Split(line, ",")
		order := []int{}
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}

			i, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			order = append(order, i)
		}

		ordering = append(ordering, order)
	}

	return ordering
}

func Day05_Part1(lines []string) {
	rules := Day05_GetRules(lines)
	sum := 0

	for _, order := range Day05_GetOrders(lines) {
		if len(order) == 0 {
			continue
		}

		valid, page := Day05_ValidRule(rules, order)
		if !valid {
			fmt.Printf("\x1b[91;1m%d: %d in bad order\x1b[0m\n", order, page)
		} else {
			fmt.Printf("\x1b[94;1m%d: valid\x1b[0m\n", order)
			sum += order[len(order)/2]
		}
	}
	fmt.Printf("Sum of middle pages: %d\n", sum)
}

func Day05_Part2(lines []string) {
	rules := Day05_GetRules(lines)
	sum := 0

	for _, order := range Day05_GetOrders(lines) {
		parents := map[int][]int{}
		valid, _ := Day05_ValidRule(rules, order)
		if valid {
			continue
		}

		fmt.Printf("Bad order: %d\n", order)

		for _, page := range order {
			for parent, children := range rules {
				if slices.Contains(order, parent) && slices.Contains(children, page) {
					parents[parent] = append(parents[parent], page)
				}
			}
		}

		fmt.Printf("Parents: %d\n", parents)

		new_order := []int{}
		for len(parents) > 0 {
			order = []int{}
			max_parent := 0
			for parent, children := range parents {
				if len(children) > len(order) {
					order = children[:]
					max_parent = parent
				}
			}

			delete(parents, max_parent)
			new_order = append(new_order, max_parent)
			if len(order) == 1 {
				new_order = append(new_order, order...)
			}

		}

		fmt.Printf("New order: %d\n\n", new_order)

		sum += new_order[len(new_order)/2]
	}

	fmt.Printf("Sum of middle pages: %d\n", sum)
}

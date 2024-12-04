package calendar

import "fmt"

func Day04_Part1(lines []string) {
	chars := [][]rune{}
	for _, line := range lines {
		chars = append(chars, []rune(line))
	}

	matches := 0

	for y := 0; y < len(chars); y++ {
		for x := 0; x < len(chars[y]); x++ {
			if chars[y][x] != 'X' {
				continue
			}

			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					if Day04_XmasDirection(chars, 'M', x+dx, y+dy) &&
						Day04_XmasDirection(chars, 'A', x+(dx*2), y+(dy*2)) &&
						Day04_XmasDirection(chars, 'S', x+(dx*3), y+(dy*3)) {
						matches++
					}
				}
			}
		}
	}

	fmt.Println(matches)
}

func Day04_XmasDirection(chars [][]rune, char rune, x int, y int) bool {
	if x < 0 || y < 0 || y >= len(chars) || x >= len(chars[y]) {
		return false
	}

	return chars[y][x] == char
}

func Day04_Part2(lines []string) {
	chars := [][]rune{}
	for _, line := range lines {
		chars = append(chars, []rune(line))
	}

	matches := 0
	// inset by one because X cannot be at the edge
	for y := 1; y < len(chars)-1; y++ {
		for x := 1; x < len(chars[y])-1; x++ {
			if chars[y][x] != 'A' {
				continue
			}

			south_west := chars[y-1][x-1] == 'M' && chars[y+1][x+1] == 'S'
			south_east := chars[y-1][x+1] == 'M' && chars[y+1][x-1] == 'S'
			north_west := chars[y+1][x-1] == 'M' && chars[y-1][x+1] == 'S'
			north_east := chars[y+1][x+1] == 'M' && chars[y-1][x-1] == 'S'

			if north_east && south_east {
				matches++
			}

			if north_east && north_west {
				matches++
			}

			if south_west && south_east {
				matches++
			}

			if south_west && north_west {
				matches++
			}
		}
	}

	fmt.Println(matches)
}

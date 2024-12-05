package calendar

import "fmt"

var around []Vector[int] = []Vector[int]{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func Day04_Part1(lines []string) {
	grid := Grid{}.from(lines)

	matches := Counter(0)

	grid.for_rune('X', func(g Grid) {
		for _, vec := range around {
			matches.Inc(g.linev(vec, "XMAS"))
		}
	})

	fmt.Println(matches)
}

func Day04_Part2(lines []string) {
	grid := Grid{}.from(lines)

	matches := Counter(0)
	grid.for_rune('A', func(grid Grid) {
		south_west := grid.d(-1, -1) == 'M' && grid.d(1, 1) == 'S'
		south_east := grid.d(1, -1) == 'M' && grid.d(-1, 1) == 'S'
		north_west := grid.d(-1, 1) == 'M' && grid.d(1, -1) == 'S'
		north_east := grid.d(1, 1) == 'M' && grid.d(-1, -1) == 'S'

		matches.Inc(north_east && south_east)
		matches.Inc(north_east && north_west)
		matches.Inc(south_west && south_east)
		matches.Inc(south_west && north_west)
	})

	fmt.Println(matches)
}

package calendar

type Counter int

func (c *Counter) Inc(yes bool) {
	if yes {
		*c++
	}
}

type Grid struct {
	runes [][]rune
	xy    Vector[int]
}

func (g Grid) from(lines []string) Grid {
	g.xy = Vector[int]{0, 0}
	g.runes = [][]rune{}
	for _, line := range lines {
		g.runes = append(g.runes, []rune(line))
	}
	return g
}

func (g *Grid) set(x, y int) rune {
	g.xy = Vector[int]{x, y}
	return g.d(0, 0)
}

func (g Grid) d(x, y int) rune {
	return g.dv(Vector[int]{x, y})
}

func (g Grid) dv(v Vector[int]) rune {
	v = g.xy.Add(v)
	if v.x < 0 || v.y < 0 || v.y >= len(g.runes) || v.x >= len(g.runes[v.y]) {
		return '\x00'
	}
	return g.runes[v.y][v.x]
}

func (g Grid) line(dx int, dy int, word string) bool {
	return g.linev(Vector[int]{dx, dy}, word)
}

func (g Grid) linev(v Vector[int], word string) bool {
	for n, r := range word {
		if g.dv(v.Mul(n)) != r {
			return false
		}
	}
	return true
}

func (g Grid) for_rune(r rune, cb func(Grid)) {
	for y := 0; y < len(g.runes); y++ {
		for x := 0; x < len(g.runes[y]); x++ {
			if g.set(x, y) == r {
				cb(g)
			}
		}
	}
}

type Vector[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64] struct {
	x, y T
}

func (v Vector[T]) Add(o Vector[T]) Vector[T] {
	return Vector[T]{v.x + o.x, v.y + o.y}
}

func (v Vector[T]) Sub(o Vector[T]) Vector[T] {
	return Vector[T]{v.x - o.x, v.y - o.y}
}

func (v Vector[T]) Mul(d T) Vector[T] {
	return Vector[T]{v.x * d, v.y * d}
}

func (v Vector[T]) Copy() Vector[T] {
	return Vector[T]{v.x, v.y}
}

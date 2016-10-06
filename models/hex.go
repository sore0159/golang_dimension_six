package models

type Hex struct {
	X int
	Y int
}

func (c1 Hex) Add(c2 Hex) (c3 Hex) {
	return Hex{X: c1.X + c2.X, Y: c1.Y + c2.Y}
}

func (c1 Hex) Subtract(c2 Hex) (c3 Hex) {
	return Hex{X: c1.X - c2.X, Y: c1.Y - c2.Y}
}

func (c1 Hex) Scale(n int) (c2 Hex) {
	return Hex{X: c1.X * n, Y: c1.Y * n}
}

func (c1 Hex) StepsTo(c2 Hex) int {
	return HexSteps(c1.X, c1.Y, c2.X, c2.Y)
}
func (c1 Hex) IsAdjacent(c2 Hex) bool {
	return HexSteps(c1.X, c1.Y, c2.X, c2.Y) == 1
}

func (h1 Hex) PathTo(h2 Hex) []Hex {
	stepsI := h1.StepsTo(h2)
	if stepsI == 0 {
		return []Hex{h1}
	} else if stepsI == 1 {
		return []Hex{h1, h2}
	}
	path := []Hex{h1}
	parts := float64(stepsI)

	c1x, c1y := HexCenter(h1.X, h1.Y)
	c2x, c2y := HexCenter(h2.X, h2.Y)
	dx := (c2x - c1x) / parts
	dy := (c2y - c1y) / parts

	for i := 1.0; i < parts; i += 1 {
		cx, cy := c1x+dx*i, c1y+dy*i
		x, y := HexAt(cx, cy)
		path = append(path, Hex{X: x, Y: y})
	}
	path = append(path, h2)
	return path
}

func (c Hex) Ring(r int) []Hex {
	if r < 0 {
		return nil
	} else if r == 0 {
		return []Hex{c}
	}
	var hexDirs = [6]Hex{
		Hex{1, 0},
		Hex{0, 1},
		Hex{-1, 1},
		Hex{-1, 0},
		Hex{0, -1},
		Hex{1, -1},
	}
	path := make([]Hex, 0, r*6)
	for i := 0; i < 6; i++ {
		leg := hexDirs[i]
		leg = c.Add(leg.Scale(r))
		legDir := hexDirs[(i+2)%6]
		for j := 0; j < r; j++ {
			path = append(path, leg.Add(legDir.Scale(j)))
		}
	}
	return path
}

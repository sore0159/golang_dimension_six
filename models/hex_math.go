// These functions assume (1,0) (0,1) and (-1,1) as the three hex directions
// Using (1,1) instead of (-1,1) has different math
package models

import (
	"math"
)

func HexSteps(x1, y1, x2, y2 int) int {
	c1z := -x1 - y1
	c2z := -x2 - y2
	dx := Abs(x2 - x1)
	dy := Abs(y2 - y1)
	dz := Abs(c2z - c1z)
	return (dx + dy + dz) / 2
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return x * -1
}

func HexAt(x, y float64) (xH, yH int) {
	var hexTriH = 0.86602540378
	x += .5
	y += hexTriH
	hX := math.Floor(x / 1.5)
	hY := math.Floor(0.5 * ((y / hexTriH) - hX))
	boxX := x - 1.5*hX
	if boxX < 1 {
		return int(hX), int(hY)
	}
	boxX -= 1
	var boxY = y - (2*hY+hX)*hexTriH
	var slope = 2 * hexTriH
	if boxY > hexTriH {
		boxY -= 2 * hexTriH
		if boxX*-slope < boxY {
			hX += 1
			return int(hX), int(hY)
		}
		return int(hX), int(hY)
	}
	if boxX*slope > boxY {
		hX += 1
		hY -= 1
		return int(hX), int(hY)
	}
	return int(hX), int(hY)
}

func HexCenter(x, y int) (xF, yF float64) {
	x2, y2 := float64(x), float64(y)
	xF = 1.5 * x2
	yF = (2*y2 + x2) * 0.86602540378
	return xF, yF
}

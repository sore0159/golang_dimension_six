package models

import (
	"math/rand"
	"sort"
)

type Die struct {
	Color int
	Size  int
	Value int
}

func RollDie(color, size int) Die {
	return Die{Color: color, Size: size, Value: rand.Intn(size) + 1}
}

func (f *Frame) CreateDiePool(r int, rocket bool) []Die {
	dice := []Die{}
	for i := 0; i < f.Basics; i += 1 {
		dice = append(dice, RollDie(COLOR_WHITE, 6))
	}
	//if rocket && f.Rockets > 0 && r == 1 {
	if rocket {
		dice = append(dice, RollDie(COLOR_RED, 8))
	}
	var hasRanged, redAlready bool
	for _, s := range f.Systems {
		c := s.Color()
		size := 6
		if c == COLOR_RED {
			if (s == MELEE && r != 0) ||
				(s == DIRECT_FIRE && r != 1) ||
				(s == ARTILLERY && r != 2) {
				continue
			}
			if redAlready {
				size = 8
			} else {
				redAlready = true
			}
		}
		dice = append(dice, RollDie(c, size))
	}
	if !hasRanged {
		dice = append(dice, RollDie(COLOR_GREEN, 8))
	}
	SortDice(dice)
	return dice
}

type Dice []Die

func (d Dice) Len() int {
	return len(d)
}
func (d Dice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d Dice) Less(i, j int) bool {
	di, dj := d[i], d[j]
	if di.Color == dj.Color {
		if di.Value == dj.Value {
			return di.Size > dj.Size
		}
		return di.Value > dj.Value
	}
	return di.Color < dj.Color
}
func SortDice(dice []Die) {
	sort.Sort(Dice(dice))
}
func (d *Die) Reroll() {
	d.Value = rand.Intn(d.Size) + 1
}

func HasDie(dice []Die, color, val int) (int, bool) {
	for i, d := range dice {
		if d.Color == color && d.Value == val {
			return i, true
		}
	}
	return -1, false
}

func DropDie(dice []Die, index int) []Die {
	return append(dice[:index], dice[index+1:]...)
}

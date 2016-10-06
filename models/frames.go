package models

type Frame struct {
	GID  int
	CID  int
	Name string
	ID   int

	Basics  int
	Rockets int
	Systems []System

	Location Hex
	DicePool []Die
	Target   [3]int // CID, FID, Range(0,1,2)
	Defense  int
	Spot     int
}

func (f *Frame) SpendRocket() {
	f.Rockets -= 1
}
func (f *Frame) SetDicePool(dice []Die) {
	f.DicePool = dice
}
func (f *Frame) ClearDicePool() {
	f.DicePool = nil
}

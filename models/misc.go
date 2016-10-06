package models

type Game struct {
	GID       int
	Skirmish  bool
	DoomClock int
}

func (g Game) InProgress() bool {
	return g.DoomClock < 12
}
func (g *Game) StartDoomClock() {
	g.DoomClock = 11
}
func (g *Game) TicDoomClock() (doom bool) {
	g.DoomClock -= 1
	return g.DoomClock < 1
}

type Company struct {
	GID  int
	CID  int
	Name string
	PPA  int
}

type Obstacle struct {
	GID        int
	Location   Hex
	Durability int
}

type Station struct {
	GID        int
	Location   []Hex
	Controller int // CID
}

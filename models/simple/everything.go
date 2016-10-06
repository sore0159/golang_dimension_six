package simple

import (
	"mule/dim_six/models"
)

type Everything struct {
	Game      models.Game
	Companies []models.Company
	Map       models.Map
	Stations  []models.Station
	Obstacles []models.Obstacle
	Frames    []models.Frame
}

func NewEverything(gid int) *Everything {
	e := &Everything{
		Game: models.Game{
			GID:       gid,
			DoomClock: 12,
		},
	}
	return e
}

func (e *Everything) GetCompany(id int) (*models.Company, bool) {
	id -= 1
	if id < 0 || id >= len(e.Companies) {
		return nil, false
	}
	return &(e.Companies[id]), true
}

func (e *Everything) ValidCompany(id int) bool {
	pnum := len(e.Companies)
	if pnum < 2 || pnum > 5 {
		return false
	}
	var frames int
	var rockets int
	for _, f := range e.Frames {
		if f.CID == id {
			frames += 1
			rockets += f.Rockets
		}
	}
	if rockets != 3 {
		return false
	}
	var min, max int
	if e.Game.Skirmish {
		switch pnum {
		case 2:
			min, max = 4, 6
		case 3:
			min, max = 3, 5
		default:
			min, max = 3, 4
		}
	} else {
		switch pnum {
		case 2:
			min, max = 5, 8
		case 3:
			min, max = 4, 7
		case 4:
			min, max = 4, 6
		default:
			min, max = 3, 5
		}
	}
	return frames >= min && frames <= max
}

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

func (e *Everything) AddCompany(name string) int {
	id := len(e.Companies) + 1
	c := models.Company{GID: e.Game.GID, CID: id, Name: name}
	e.Companies = append(e.Companies, c)
	return id
}
func (e *Everything) GetCompany(id int) (*models.Company, bool) {
	id -= 1
	if id < 0 || id >= len(e.Companies) {
		return nil, false
	}
	return &(e.Companies[id]), true
}

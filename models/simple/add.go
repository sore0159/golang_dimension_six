package simple

import (
	"mule/dim_six/models"
)

func (e *Everything) AddCompany(name string) int {
	id := len(e.Companies) + 1
	c := models.Company{GID: e.Game.GID, CID: id, Name: name}
	e.Companies = append(e.Companies, c)
	return id
}

func (e *Everything) AddFrame(cid int, name string, rockets int, systems ...models.System) int {
	id := 1
	for _, f := range e.Frames {
		if f.CID == cid && f.FID >= id {
			id = f.FID + 1
		}
	}
	f := models.Frame{
		GID:  e.Game.GID,
		CID:  cid,
		FID:  id,
		Name: name,

		Basics:  2,
		Rockets: rockets,
		Systems: systems,
	}
	e.Frames = append(e.Frames, f)
	return id
}

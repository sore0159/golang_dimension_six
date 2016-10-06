package main

import (
	"mule/dim_six/models/simple"
)

var GAMEDATA []*simple.Everything

func GetData(gid int) (*simple.Everything, bool) {
	gid -= 1
	if gid < 0 || gid >= len(GAMEDATA) {
		return nil, false
	}
	return GAMEDATA[gid], true
}

func NewData() int {
	id := len(GAMEDATA) + 1
	e := simple.NewEverything(id)
	GAMEDATA = append(GAMEDATA, e)
	return id
}

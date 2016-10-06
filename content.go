package main

import (
	"net/http"
)

func pageMainIndex(w http.ResponseWriter, r *http.Request) {
	numGames := len(GAMEDATA)
	gameList := make([]int, numGames)
	for i, _ := range gameList {
		gameList[i] = i + 1
	}
	serveTemplate(w, gameList, "mainIndex")
}

func pageGame(w http.ResponseWriter, r *http.Request) {
	ids, oks := GetInts(r.URL.Path, 2)
	if !oks[0] {
		http.Error(w, "Unparsable gid", 500)
		return
	}
	id := ids[0]
	e, ok := GetData(id)
	if !ok {
		http.Error(w, "gid not found", 500)
		return
	}
	pgd := pageGameData{
		GID:       e.Game.GID,
		Doom:      e.Game.DoomClock,
		CanAdd:    !e.Game.InProgress() && len(e.Companies) < 5,
		Companies: []string{""},
	}
	if !e.Game.InProgress() && len(e.Companies) > 1 {
		pgd.CanStart = true
	}
	for id := 1; ; id += 1 {
		c, ok := e.GetCompany(id)
		if !ok {
			break
		}
		pgd.Companies = append(pgd.Companies, c.Name)
	}
	serveTemplate(w, pgd, "gameIndex")
}

type pageGameData struct {
	GID       int
	Doom      int
	CanAdd    bool
	CanStart  bool
	Companies []string
}

/*
func pageCompany(w http.ResponseWriter, r *http.Request) {
	ids, oks := GetInts(r.URL.Path, 2, 3)
	if !oks[0] || !ok[1] {
		http.Error(w, "Unparsable gid/cid", 500)
		return
	}
	gid, cid := ids[0], ids[1]
	e, ok := GetData(id)
	if !ok {
		http.Error(w, "gid not found", 500)
		return
	}
	pgd := pageCompanyData{
		GID:       e.Game.GID,
		Doom:      e.Game.DoomClock,
		CanAdd:    !e.Game.InProgress() && len(e.Companies) < 5,
		Companies: []string{""},
	}
	if !e.Game.InProgress() && len(e.Companies) > 1 {
		pgd.CanStart = true
	}
	for id := 1; ; id += 1 {
		c, ok := e.GetCompany(id)
		if !ok {
			break
		}
		pgd.Companies = append(pgd.Companies, c.Name)
	}
	serveTemplate(w, pgd, "gameIndex")
}

type pageCompanyData struct {
	GID    int
	CID    int
	Frames []frameData
}
type frameData struct {
	FID     int
	Name    string
	Systems []string
	Rockets int
}
*/

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func pageMainIndex(w http.ResponseWriter, r *http.Request) {
	numGames := len(GAMEDATA)
	gameList := make([]int, numGames)
	for i, _ := range gameList {
		gameList[i] = i + 1
	}
	serveTemplate(w, gameList, "mainIndex")
}

func commandNewGame(w http.ResponseWriter, r *http.Request) {
	id := NewData()
	http.Redirect(w, r, fmt.Sprintf("/game/%d", id), http.StatusFound)
}

func commandNewCompany(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/newcompany/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Path[12:])
	if err != nil {
		http.Error(w, "Unparsable gid "+r.URL.Path[12:], 500)
		return
	}
	e, ok := GetData(id)
	if !ok {
		http.Error(w, "gid not found", 500)
		return
	}
	if len(e.Companies) > 4 {
		http.Error(w, "max 5 companies", 500)
		return
	}
	if e.Game.InProgress() {
		http.Error(w, "cannot add company mid-game", 500)
		return
	}
	e.AddCompany("TEST COMPANY")
	http.Redirect(w, r, fmt.Sprintf("/game/%d", id), http.StatusFound)
}

func pageGame(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/game/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Path[6:])
	if err != nil {
		http.Error(w, "Unparsable gid", 500)
		return
	}
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

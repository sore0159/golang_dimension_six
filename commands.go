package main

/*
import (
	"fmt"
	"mule/dim_six/models"
	"net/http"
)

func commandNewGame(w http.ResponseWriter, r *http.Request) {
//	id := NewData()
	http.Redirect(w, r, fmt.Sprintf("/game/%d", id), http.StatusFound)
}

func commandNewCompany(w http.ResponseWriter, r *http.Request) {
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

func commandNewFrame(w http.ResponseWriter, r *http.Request) {
	ids, oks := GetInts(r.URL.Path, 2, 3)
	if !oks[0] || !oks[1] {
		http.Error(w, "Unparsable gid/cid", 500)
		return
	}
	gid, cid := ids[0], ids[1]
	e, ok := GetData(gid)
	if !ok {
		http.Error(w, "gid not found", 500)
		return
	}
	if e.Game.InProgress() {
		http.Error(w, "cannot add frame mid-game", 500)
		return
	}
	e.AddFrame(cid, "TEST FRAME", 1, models.DEFENSE)
	http.Redirect(w, r, fmt.Sprintf("/company/%d/%d", gid, cid), http.StatusFound)
}
*/

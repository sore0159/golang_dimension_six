package main

import (
	"github.com/julienschmidt/httprouter"
	er "mule/dim_six/errors"
	"net/http"
)

func jsonGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	NewData()
	list := make([]int, 0, len(GAMEDATA))
	for _, e := range GAMEDATA {
		list = append(list, e.Game.GID)
	}
	JSONSuccess(w, list)
}

func jsonGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids, ok := AtoI(ps[0].Value)
	if !ok {
		JSONUserError(w, er.NewU("Unparsable GID"))
		return
	}
	e, ok := GetData(ids[0])
	if !ok {
		JSONUserError(w, er.NewU("GID Not Found"))
		return
	}
	JSONSuccess(w, e.Game)
}

func jsonCompanies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids, ok := AtoI(ps[0].Value)
	if !ok {
		JSONUserError(w, er.NewU("Unparsable GID"))
		return
	}
	e, ok := GetData(ids[0])
	if ok {
		JSONSuccess(w, e.Companies)
		return
	}
	JSONUserError(w, er.NewU("GID Not Found"))
}

func jsonCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids, ok := AtoI(ps[0].Value, ps[1].Value)
	if !ok {
		JSONUserError(w, er.NewU("Unparsable GID/CID"))
		return
	}
	e, ok := GetData(ids[0])
	if ok {
		if c, ok := e.GetCompany(ids[1]); ok {
			JSONSuccess(w, c)
			return
		}
	}
	JSONUserError(w, er.NewU("GID+CID Not Found"))
}

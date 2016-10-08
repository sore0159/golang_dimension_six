package main

import (
	"github.com/julienschmidt/httprouter"
	er "mule/dim_six/errors"
	"net/http"
)

func pageMainIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serveTemplate(w, nil, "mainIndex")
}

// link/create games
func pageGameIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serveTemplate(w, nil, "gameIndex")
}

// create/link companies, config game, start game
func pageGameSetup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids, ok := AtoI(ps[0].Value)
	if !ok {
		HTMLUserError(w, er.NewU("Unparsable GID"))
		return
	}
	serveTemplate(w, ids[0], "gameSetup")
}

// link companies, game info
func pageGameView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids, ok := AtoI(ps[0].Value)
	if !ok {
		HTMLUserError(w, er.NewU("Unparsable GID"))
		return
	}
	serveTemplate(w, ids[0], "gameView")
}

func pageCompanySetup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids, ok := AtoI(ps[0].Value, ps[1].Value)
	if !ok {
		HTMLUserError(w, er.NewU("Unparsable GID/CID"))
		return
	}
	serveTemplate(w, ids, "companySetup")
}

func pageCompanyView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids, ok := AtoI(ps[0].Value, ps[1].Value)
	if !ok {
		HTMLUserError(w, er.NewU("Unparsable GID/CID"))
		return
	}
	serveTemplate(w, ids, "companyView")
}

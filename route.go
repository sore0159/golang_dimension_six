package main

import (
	"github.com/julienschmidt/httprouter"
	//	er "mule/dim_six/errors"
	"net/http"
	"strconv"
)

func MakeRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/favicon.ico", imgFavIcon)
	router.ServeFiles("/static/*filepath", http.Dir("FILES/static/"))

	router.GET("/", pageMainIndex)

	router.GET("/games", pageGameIndex)
	router.GET("/games/:gid", pageGameView)
	router.GET("/games/:gid/setup", pageGameSetup)
	router.GET("/json/games", jsonGames)
	router.GET("/json/games/:gid", jsonGame)

	router.GET("/companies/:gid/:cid", pageCompanyView)
	router.GET("/companies/:gid/:cid/setup", pageCompanySetup)
	router.GET("/json/companies/:gid", jsonCompanies)
	router.GET("/json/companies/:gid/:cid", jsonCompany)
	return router
}

func imgFavIcon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "FILES/static/img/favicon.ico")
}

func AtoI(strs ...string) ([]int, bool) {
	r := make([]int, len(strs))
	var err error
	for i, str := range strs {
		if r[i], err = strconv.Atoi(str); err != nil {
			return nil, false
		}
	}
	return r, true
}

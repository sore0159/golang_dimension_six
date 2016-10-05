package main

import (
	"net/http"
)

func SetupMux() {
	http.HandleFunc("/", pageMainIndex)
	http.HandleFunc("/favicon.ico", imgFavIcon)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("FILES/static/"))))
	http.HandleFunc("/json", apiJSON)
}

func imgFavIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "FILES/static/img/favicon.ico")
}

func pageMainIndex(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, nil, "mainIndex")
}

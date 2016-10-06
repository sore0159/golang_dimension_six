package main

import (
	"net/http"
	"strconv"
	"strings"
)

func SetupMux() {
	http.HandleFunc("/", pageMainIndex)
	http.HandleFunc("/favicon.ico", imgFavIcon)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("FILES/static/"))))

	http.HandleFunc("/json", apiJSON)

	http.HandleFunc("/game/", pageGame)

	http.HandleFunc("/newgame", commandNewGame)
	http.HandleFunc("/newcompany/", commandNewCompany)
	http.HandleFunc("/newframe/", commandNewFrame)

}

func imgFavIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "FILES/static/img/favicon.ico")
}

func GetInts(path string, iList ...int) ([]int, []bool) {
	parts := strings.Split(path, "/")
	res, ok := make([]int, len(iList)), make([]bool, len(iList))
	for j, i := range iList {
		if i >= len(parts) {
			res[j], ok[j] = 0, false
			continue
		}
		x, err := strconv.Atoi(parts[i])
		if err != nil {
			res[j], ok[j] = 0, false
			continue
		}
		res[j], ok[j] = x, true
	}
	return res, ok
}

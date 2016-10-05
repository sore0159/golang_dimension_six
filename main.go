package main

import (
	//"fmt"
	"log"
	er "mule/dim_six/errors"
	"net/http"
)

func main() {
	port := ":8080"
	log.Println("STARTING SERVER ON PORT " + port)
	SetupMux()
	if s := er.Check(http.ListenAndServe(port, nil)); s != nil {
		log.Println("Listen And Serve Error")
		Log(s)
	}
}

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
	serveTemplate(w, nil, "hello")
	//	fmt.Fprint(w, "HELLO WORLD")
}

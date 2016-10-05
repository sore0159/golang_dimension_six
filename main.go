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

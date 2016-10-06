package main

import (
	//"fmt"
	"log"
	er "mule/dim_six/errors"
	lg "mule/dim_six/log"
	//	"mule/dim_six/views"
	"net/http"
)

func main() {
	port := ":8080"
	log.Println("STARTING SERVER ON PORT " + port)
	SetupMux()
	if err := http.ListenAndServe(port, nil); err != nil {
		lg.Log(er.NewS(err, "main ListenAndServe failure"))
	}
}

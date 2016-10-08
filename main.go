package main

import (
	//"fmt"
	"log"
	er "mule/dim_six/errors"
	//	"mule/dim_six/views"
	"net/http"
)

func main() {
	port := ":8080"
	log.Println("STARTING SERVER ON PORT " + port)
	if err := http.ListenAndServe(port, MakeRouter()); err != nil {
		Log(er.NewS(err, "main ListenAndServe failure"))
	}
}

package main

import (
	er "mule/dim_six/errors"
	"net/http"
)

func apiJSONget(w http.ResponseWriter, r *http.Request) {
	JSONServerError(w, er.NewS(er.NewU("TODO: JSON GET"), "TODO"))
}
func apiJSONput(w http.ResponseWriter, r *http.Request) {
	JSONUserError(w, er.NewU("TODO: JSON PUT"))
}

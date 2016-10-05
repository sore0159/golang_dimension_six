package main

import (
	"html/template"
	er "mule/dim_six/errors"
	"net/http"
)

func serveTemplate(w http.ResponseWriter, data interface{}, files ...string) {
	for i, n := range files {
		files[i] = "FILES/templates/" + n + ".html"
	}
	t, err := template.ParseFiles(files...)
	if s := er.Check(err); s != nil {
		Log(s)
		return
	}
	err = t.Execute(w, data)
	if s := er.Check(err); s != nil {
		Log(s)
		return
	}
}

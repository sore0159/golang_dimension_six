package main

import (
	"html/template"
	er "mule/dim_six/errors"
	"net/http"
)

func serveTemplate(w http.ResponseWriter, data interface{}, files ...string) {
	if len(files) == 0 {
		w.WriteHeader(500)
		w.Write([]byte("error: no html templates given"))
		return
	}
	if files[0] != "frame" {
		files = append([]string{"frame"}, files...)
	}
	for i, n := range files {
		files[i] = "FILES/templates/" + n + ".html"
	}
	t, err := template.ParseFiles(files...)
	if s := er.Check(err); s != nil {
		Log(s)
		return
	}
	err = t.ExecuteTemplate(w, "frame", data)
	if s := er.Check(err); s != nil {
		Log(s)
		return
	}
}

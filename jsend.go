package main

import (
	"encoding/json"
	er "mule/dim_six/errors"
	"net/http"
)

func apiJSON(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		apiJSONget(w, r)
	case "PUT":
		apiJSONput(w, r)
	default:
		JSONUserError(w, er.NewU("UNSUPPORTED JSON METHOD"))
	}
}

func JSONSuccess(w http.ResponseWriter, obj interface{}) {
	raw, err := json.Marshal(obj)
	if c := er.Check(err); c != nil {
		JSONServerError(w, c)
		return
	}
	data := []byte("{'status':'success', 'data': ")
	data = append(data, raw...)
	data = append(data, []byte(`}`)...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if c := er.Check(err); c != nil {
		Log(c)
	}
}

func JSONUserError(w http.ResponseWriter, u *er.User) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(400)
	_, err := w.Write([]byte("{'status': 'fail', 'data': {'message':'" + u.Error() + "'}}"))
	if c := er.Check(err); c != nil {
		Log(c)
	}
}

func JSONServerError(w http.ResponseWriter, s *er.Server) {
	Log(s)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("{'status': 'error', 'message': 'internal server error'}"))
	if c := er.Check(err); c != nil {
		Log(c)
	}
}

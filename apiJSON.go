package main

import (
	"encoding/json"
	er "mule/dim_six/errors"
	"net/http"
)

func JSONSuccess(w http.ResponseWriter, obj interface{}) {
	raw, err := json.Marshal(obj)
	if err != nil {
		JSONServerError(w, er.NewS(err, "json success marshal failure"))
		return

	}
	data := []byte("{'status':'success', 'data': ")
	data = append(data, raw...)
	data = append(data, []byte(`}`)...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(data); err != nil {
		Log(er.NewS(err, "json success write failure"))
	}
}

func JSONUserError(w http.ResponseWriter, u *er.User) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(400)
	_, err := w.Write([]byte("{'status': 'fail', 'data': {'message':'" + u.Error() + "'}}"))
	if err != nil {
		Log(er.NewS(err, "json user error write failure"))
	}
}

func JSONServerError(w http.ResponseWriter, s *er.Server) {
	Log(s)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("{'status': 'error', 'message': 'internal server error'}"))
	if err != nil {
		Log(er.NewS(err, "json server error write failure"))
	}
}

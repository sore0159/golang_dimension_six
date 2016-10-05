package errors

import (
	"runtime"
)

type User string

func NewU(msg string) *User {
	u := User(msg)
	return &u
}

func (u *User) Error() string {
	if u == nil {
		return "NIL USER ERROR"
	}
	msg := string(*u)
	if msg == "" {
		return "BLANK USER ERROR"
	}
	return msg
}

type Server struct {
	Base   error
	Layers []*Context
}

func NewS(err error, msg string) *Server {
	_, fName, lineNum, _ := runtime.Caller(1)
	c := &Context{File: fName, Line: lineNum,
		DataKeys: []string{"creation"},
		DataVals: []interface{}{msg},
	}
	return &Server{
		Base:   err,
		Layers: []*Context{c},
	}
}

func (s *Server) Error() string {
	return "Server Error, Base: " + s.Base.Error()
}

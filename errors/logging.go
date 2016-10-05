package errors

import (
	"fmt"
	"path"
)

func (s *Server) Log(logger func(v ...interface{})) {
	logger("Server Error, Base: " + s.Base.Error())
}

func (s *Server) Verbose(logger func(v ...interface{})) {
	msg := "Server Error, Base: " + s.Base.Error()
	for i, c := range s.Layers {
		fileStr := path.Base(c.File)
		msg += fmt.Sprintf("\nContext(%d) %s(%02d):", i, fileStr, c.Line)
		for i, key := range c.DataKeys {
			val := c.DataVals[i]
			msg += fmt.Sprintf("\n%s : %v", key, val)
		}
		msg += "\n"
	}
	logger(msg)
}

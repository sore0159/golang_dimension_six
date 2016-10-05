package errors

import (
	"runtime"
)

type Context struct {
	File     string
	Line     int
	DataKeys []string
	DataVals []interface{}
}

func (s *Server) NewContext(key string, val interface{}) *Context {
	_, fName, lineNum, _ := runtime.Caller(1)
	c := &Context{File: fName, Line: lineNum,
		DataKeys: []string{key},
		DataVals: []interface{}{val},
	}
	s.Layers = append(s.Layers, c)
	return c
}

func (c *Context) AddData(key string, val interface{}) *Context {
	c.DataKeys = append(c.DataKeys, key)
	c.DataVals = append(c.DataVals, val)
	return c
}

func (s *Server) Last() *Context {
	return s.Layers[len(s.Layers)-1]
}

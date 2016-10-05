package errors

type User string

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

func NewU(msg string) *User {
	u := User(msg)
	return &u
}

type Server struct {
	Base   error
	Layers []*Layer
}

func Check(err error) *Server {
	if err == nil {
		return nil
	}
	return &Server{Base: err}
}

type Layer struct {
	File string
	Line int
	Data map[string]interface{}
}

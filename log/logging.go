package log

import (
	"log"
	er "mule/dim_six/errors"
)

func Log(s *er.Server) {
	s.Verbose(log.Println)
}

package errors

import (
	"errors"
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	err := errors.New("Test Error")
	s := NewS(err, "Test Wrapping")
	s.Last().AddData("ctest", 3)
	s.NewContext("Context Test", 1).AddData("More Data", 2)
	s.Verbose(log.Println)
}

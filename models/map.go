package models

import (
	"fmt"
)

type Map struct {
	GID   int
	Areas []Area
}

func (m Map) Valid(h Hex) bool {
	for _, a := range m.Areas {
		if a.Valid(h) {
			return true
		}
	}
	return false
}

type Area interface {
	Valid(Hex) bool
	MarshalJSON() ([]byte, error)
}

type AreaBurst struct {
	Center Hex
	Size   int
}

func (a AreaBurst) Valid(h Hex) bool {
	return h.StepsTo(a.Center) <= a.Size
}

func (a AreaBurst) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{'center': {'x':%d, 'y':%d}, 'size':%d }", a.Center.X, a.Center.Y, a.Size)), nil
}

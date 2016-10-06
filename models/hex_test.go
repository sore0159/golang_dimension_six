package models

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	h1 := Hex{0, 0}
	h2 := Hex{-5, 0}
	path := h1.PathTo(h2)
	steps := h1.StepsTo(h2)
	log.Println("STEPS, PATH:", steps, path)
}

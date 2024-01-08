package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Actor interface {
	Allegiance() []Allegiance
	Draw(*ebiten.Image)
	Position() (uint, uint)
	Update(Room) error
}

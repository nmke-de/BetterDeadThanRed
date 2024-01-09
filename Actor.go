package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// (interface for player character, mobs, NPCs, perhaps even obstacles etc)
type Actor interface {
	Alive() bool
	Allegiance() []Allegiance
	Draw(*ebiten.Image)
	Position() (uint, uint)
	Update(Room) error
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// (interface for player character, mobs, NPCs, perhaps even obstacles etc)
type Actor interface {
	Alive() bool
	Allegiance() []Allegiance
	Collide(Actor)
	Draw(*ebiten.Image)
	Hitbox() uint
	Position() (uint, uint)
	TakeDamage(int)
	Update() error
}

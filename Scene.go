package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Draw(*ebiten.Image, *Game)
	Update(*Game, []ebiten.Key) error
}

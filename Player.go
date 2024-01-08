package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const on_frames = 7

type Player struct {
	state *PlayerState
	imgs []*ebiten.Image
}

type PlayerState struct {
	x, y uint
	animation_state uint
}

func newPlayer(x, y uint) Player {
	return Player {
		&PlayerState {
			x, y,
			0,
		},
		[]*ebiten.Image{
			ebiten.NewImageFromImage(loadPNG("Libright.png")),
			ebiten.NewImageFromImage(loadPNG("Libright_Jump.png")),
		},
	}
}

func (p Player) Allegiance() []Allegiance {
	return []Allegiance{player}
}

func (p Player) Draw(surface *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.state.x), float64(p.state.y))
	surface.DrawImage(p.imgs[p.state.animation_state / on_frames], op)
}

func (p Player) Position() (uint, uint) {
	return p.state.x, p.state.y
}

func (p Player) Update(width, height uint) error {
	movement := false
	pressed := inpututil.AppendPressedKeys([]ebiten.Key{})
	for _, key := range pressed {
		switch key {
		case ebiten.KeyA:
			p.state.x = uint(max(int(p.state.x - 1), 0))
			movement = true
		case ebiten.KeyD:
			p.state.x = min(p.state.x + 1, width)
			movement = true
		case ebiten.KeyS:
			p.state.y = min(p.state.y + 1, height)
			movement = true
		case ebiten.KeyW:
			p.state.y = uint(max(int(p.state.y - 1), 0))
			movement = true
		}
	}
	if movement {
		p.state.animation_state = (p.state.animation_state + 1) % uint(len(p.imgs) * on_frames)
	} else {
		p.state.animation_state = 0
	}
	return nil
}

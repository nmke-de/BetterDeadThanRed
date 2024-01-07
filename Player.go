package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

type Player struct {
	state *PlayerState
	img *ebiten.Image
}

type PlayerState struct {
	x, y uint
}

func newPlayer(x, y uint) Player {
	img := ebiten.NewImage(10, 10)
	img.Fill(color.RGBA{uint8(255), uint8(128), uint8(32), 255})
	return Player {
		&PlayerState {
			x, y,
		},
		img,
	}
}

func (p Player) Draw(surface *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.state.x), float64(p.state.y))
	surface.DrawImage(p.img, op)
}

func (p Player) Update(width, height uint) error {
	pressed := inpututil.AppendPressedKeys([]ebiten.Key{})
	for _, key := range pressed {
		switch key {
		case ebiten.KeyA:
			p.state.x = uint(max(int(p.state.x - 1), 0))
		case ebiten.KeyD:
			p.state.x = min(p.state.x + 1, width)
		case ebiten.KeyS:
			p.state.y = min(p.state.y + 1, height)
		case ebiten.KeyW:
			p.state.y = uint(max(int(p.state.y - 1), 0))
		}
	}
	return nil
}

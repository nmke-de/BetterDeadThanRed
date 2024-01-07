package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

type ComradeVodka struct {
	state *ComradeVodkaState
	img *ebiten.Image
}

type ComradeVodkaState struct {
	x, y uint
}

func newComradeVodka(x, y uint) ComradeVodka {
	img := ebiten.NewImage(10, 10)
	img.Fill(color.RGBA{uint8(255), uint8(16), uint8(32), 255})
	return ComradeVodka {
		&ComradeVodkaState {
			x, y,
		},
		img,
	}
}

func (c ComradeVodka) Allegiance() []Allegiance {
	return []Allegiance{commie, comradeVodka}
}

func (c ComradeVodka) Draw(surface *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.state.x), float64(c.state.y))
	surface.DrawImage(c.img, op)
}

func (c ComradeVodka) Position() (uint, uint) {
	return c.state.x, c.state.y
}

func (c ComradeVodka) Update(width, height uint) error {
	w := int(width)
	h := int(height)
	x := 0
	if (rand.Int() % 2 == 0) {
		x = -1
	} else {
		x = 1
	}
	y := 0
	if (rand.Int() % 2 == 0) {
		y = -1
	} else {
		y = 1
	}
	c.state.x = uint(min(max(int(c.state.x) + x, 0), w))
	c.state.y = uint(min(max(int(c.state.y) + y, 0), h))
	return nil
}

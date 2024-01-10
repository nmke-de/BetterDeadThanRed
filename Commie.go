package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Commie struct {
	state *CommieState
	img   *ebiten.Image
}

type CommieState struct {
	x, y uint
}

func newCommie(x, y uint) Commie {
	img := ebiten.NewImage(10, 10)
	img.Fill(color.RGBA{uint8(255), uint8(16), uint8(32), 255})
	return Commie{
		&CommieState{
			x, y,
		},
		img,
	}
}

func (c Commie) Alive() bool {
	return true
}

func (c Commie) Allegiance() []Allegiance {
	return []Allegiance{commie}
}

func (c Commie) Collide(a Actor) {
	// TODO collision detection
	return
}

func (c Commie) Draw(surface *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.state.x), float64(c.state.y))
	surface.DrawImage(c.img, op)
}

func (c Commie) Hitbox() uint {
	return 15
}

func (c Commie) Position() (uint, uint) {
	return c.state.x, c.state.y
}

func (c Commie) Update(r Room) error {
	w := int(r.width)
	h := int(r.height)

	// Find player
	px, py := (*r.actors)[r.cache["player"]].Position()

	// Move
	if px < c.state.x {
		c.state.x--
	} else if px > c.state.x {
		c.state.x++
	}
	if py < c.state.y {
		c.state.y--
	} else if py > c.state.y {
		c.state.y++
	}
	c.state.x = uint(min(max(int(c.state.x), 0), w))
	c.state.y = uint(min(max(int(c.state.y), 0), h))
	return nil
}

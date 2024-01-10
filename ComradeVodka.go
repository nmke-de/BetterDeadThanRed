package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

type ComradeVodka struct {
	state *ComradeVodkaState
	img   *ebiten.Image
	roomname string
}

type ComradeVodkaState struct {
	x, y uint
}

func newComradeVodka(x, y uint, roomname string) ComradeVodka {
	img := ebiten.NewImage(10, 10)
	img.Fill(color.RGBA{uint8(255), uint8(16), uint8(32), 255})
	return ComradeVodka{
		&ComradeVodkaState{
			x, y,
		},
		img,
		roomname,
	}
}

func (c ComradeVodka) Alive() bool {
	return true
}

func (c ComradeVodka) Allegiance() []Allegiance {
	return []Allegiance{commie, comradeVodka}
}

func (c ComradeVodka) Collide(a Actor) {
	cx, cy := c.Position()
	ax, ay := a.Position()
	r := Room(scenes[c.roomname].(Room))
	c.state.x = clamp((cx + cx - ax), 0, r.width)
	c.state.y = clamp((cy + cy - ay), 0, r.height)
}

func (c ComradeVodka) Draw(surface *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.state.x), float64(c.state.y))
	surface.DrawImage(c.img, op)
}

func (c ComradeVodka) Hitbox() uint {
	return 15
}

func (c ComradeVodka) Position() (uint, uint) {
	return c.state.x, c.state.y
}

func (c ComradeVodka) Update(r Room) error {
	w := int(r.width)
	h := int(r.height)
	x := rand.Int()%3 - 1
	y := rand.Int()%3 - 1
	c.state.x = uint(min(max(int(c.state.x)+x, 0), w))
	c.state.y = uint(min(max(int(c.state.y)+y, 0), h))
	return nil
}

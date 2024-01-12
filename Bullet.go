package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Bullet struct {
	roomname string
	vx, vy   int
	state    *CharacterState
}

func newBullet(x, y, vx, vy int, roomname string) Bullet {
	return Bullet{
		roomname,
		vx, vy,
		&CharacterState{
			uint(x), uint(y), 0, 1,
		},
	}
}

func (b Bullet) Alive() bool {
	r := Room(scenes[b.roomname].(Room))
	return b.state.health > 0 && b.state.x > 0 && b.state.y > 0 && b.state.x < r.width && b.state.y < r.height
}

func (b Bullet) Allegiance() []Allegiance {
	return []Allegiance{bullet}
}

func (b Bullet) Collide(a Actor) {
	// TODO damage
	a.TakeDamage(2)
	b.state.health = 0
}

func (b Bullet) Draw(surface *ebiten.Image) {
	img := ebiten.NewImage(2, 2)
	img.Fill(color.RGBA{uint8(128), uint8(128), uint8(128), uint8(255)})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.state.x), float64(b.state.y))
	surface.DrawImage(img, op)
}

func (b Bullet) Hitbox() uint {
	return 2
}

func (b Bullet) Position() (uint, uint) {
	return b.state.x, b.state.y
}

func (b Bullet) TakeDamage(damage int) {
	b.state.health -= damage
}

func (b Bullet) Update() error {
	r := Room(scenes[b.roomname].(Room))
	b.state.x = clamp(b.state.x + uint(b.vx), 0, r.width)
	b.state.y = clamp(b.state.y + uint(b.vy), 0, r.height)
	return nil
}

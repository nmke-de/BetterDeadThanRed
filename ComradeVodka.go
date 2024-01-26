package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type ComradeVodka struct {
	state    *CharacterState
	imgs     []*ebiten.Image
	roomname string
}

func newComradeVodka(x, y uint, roomname string) ComradeVodka {
	return ComradeVodka{
		&CharacterState{
			x, y, 0, 2,
		},
		[]*ebiten.Image{
			ebiten.NewImageFromImage(loadPNG("Commie.png")),
			ebiten.NewImageFromImage(loadPNG("Commie_Jump.png")),
		},
		roomname,
	}
}

func (c ComradeVodka) Alive() bool {
	return c.state.health > 0
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
	surface.DrawImage(c.imgs[c.state.animation_state/on_frames], op)
}

func (c ComradeVodka) Hitbox() uint {
	return 15
}

func (c ComradeVodka) Position() (uint, uint) {
	return c.state.x, c.state.y
}

func (c ComradeVodka) TakeDamage(damage int) {
	c.state.health -= damage
}

func (c ComradeVodka) Update() error {
	r := Room(scenes[c.roomname].(Room))
	w := int(r.width)
	h := int(r.height)
	x := rand.Int()%3 - 1
	y := rand.Int()%3 - 1
	c.state.x = uint(min(max(int(c.state.x)+x, 0), w))
	c.state.y = uint(min(max(int(c.state.y)+y, 0), h))
	if abs(x)+abs(y) > 0 {
		c.state.animation_state = (c.state.animation_state + 1) % uint(len(c.imgs)*on_frames)
	} else {
		c.state.animation_state = 0
		// Shoot bullet
		bvx := rand.Int()%3 - 1
		bvy := rand.Int()%3 - 1
		if bvx != 0 || bvy != 0 {
			bullet := newBullet(int(c.state.x+c.Hitbox())+int(c.Hitbox())*bvx*2, int(c.state.y+c.Hitbox())+int(c.Hitbox())*bvy*2, bvx, bvy, c.roomname)
			*r.actors = append(*r.actors, bullet)
		}
	}
	return nil
}

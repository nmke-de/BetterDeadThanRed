package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Commie struct {
	state    *CharacterState
	imgs     []*ebiten.Image
	roomname string
}

func newCommie(x, y uint, roomname string) Commie {
	return Commie{
		&CharacterState{
			x, y, 0, 3,
		},
		[]*ebiten.Image{
			ebiten.NewImageFromImage(loadPNG("Commie.png")),
			ebiten.NewImageFromImage(loadPNG("Commie_Jump.png")),
		},
		roomname,
	}
}

func (c Commie) Alive() bool {
	return c.state.health > 0
}

func (c Commie) Allegiance() []Allegiance {
	return []Allegiance{commie}
}

func (c Commie) Collide(a Actor) {
	for _, v := range a.Allegiance() {
		if v == player {
			a.TakeDamage(1)
		}
	}
	cx, cy := c.Position()
	ax, ay := a.Position()
	r := Room(scenes[c.roomname].(Room))
	c.state.x = clamp((cx + cx - ax), 0, r.width)
	c.state.y = clamp((cy + cy - ay), 0, r.height)
}

func (c Commie) Draw(surface *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.state.x), float64(c.state.y))
	surface.DrawImage(c.imgs[c.state.animation_state/on_frames], op)
}

func (c Commie) Hitbox() uint {
	return 15
}

func (c Commie) Position() (uint, uint) {
	return c.state.x, c.state.y
}

func (c Commie) TakeDamage(dmg int) {
	c.state.health -= dmg
}

func (c Commie) Update() error {
	r := Room(scenes[c.roomname].(Room))
	w := int(r.width)
	h := int(r.height)

	// Find player
	px, py := (*r.actors)[r.cache["player"]].Position()

	// Move
	movement_x := true
	if px < c.state.x {
		c.state.x--
	} else if px > c.state.x {
		c.state.x++
	} else {
		movement_x = false
	}
	movement_y := true
	if py < c.state.y {
		c.state.y--
	} else if py > c.state.y {
		c.state.y++
	} else {
		movement_y = false
	}
	c.state.x = uint(min(max(int(c.state.x), 0), w))
	c.state.y = uint(min(max(int(c.state.y), 0), h))
	if movement_x || movement_y {
		c.state.animation_state = (c.state.animation_state + 1) % uint(len(c.imgs)*on_frames)
	} else {
		c.state.animation_state = 0
	}
	return nil
}

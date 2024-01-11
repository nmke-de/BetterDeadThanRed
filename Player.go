package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	state    *CharacterState
	imgs     []*ebiten.Image
	roomname string
}

func newPlayer(x, y uint, roomname string) Player {
	return Player{
		&CharacterState{
			x, y, 0, 5,
		},
		[]*ebiten.Image{
			ebiten.NewImageFromImage(loadPNG("Libright.png")),
			ebiten.NewImageFromImage(loadPNG("Libright_Jump.png")),
		},
		roomname,
	}
}

func (p Player) Alive() bool {
	return p.state.health > 0
}

func (p Player) Allegiance() []Allegiance {
	return []Allegiance{player}
}

func (p Player) Collide(a Actor) {
	px, py := p.Position()
	ax, ay := a.Position()
	r := Room(scenes[p.roomname].(Room))
	p.state.x = clamp((px + px - ax), 0, r.width)
	p.state.y = clamp((py + py - ay), 0, r.height)
}

func (p Player) Draw(surface *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.state.x), float64(p.state.y))
	surface.DrawImage(p.imgs[p.state.animation_state/on_frames], op)
}

func (p Player) Hitbox() uint {
	return 15
}

func (p Player) Position() (uint, uint) {
	return p.state.x, p.state.y
}

func (p Player) TakeDamage(damage int) {
	p.state.health -= damage
}

func (p Player) Update() error {
	r := Room(scenes[p.roomname].(Room))
	movement := false
	pressed := inpututil.AppendPressedKeys([]ebiten.Key{})
	for _, key := range pressed {
		switch key {
		case ebiten.KeyA:
			p.state.x = uint(max(int(p.state.x-1), 0))
			movement = true
		case ebiten.KeyD:
			p.state.x = min(p.state.x+1, r.width)
			movement = true
		case ebiten.KeyS:
			p.state.y = min(p.state.y+1, r.height)
			movement = true
		case ebiten.KeyW:
			p.state.y = uint(max(int(p.state.y-1), 0))
			movement = true
		}
	}
	if movement {
		p.state.animation_state = (p.state.animation_state + 1) % uint(len(p.imgs)*on_frames)
	} else {
		p.state.animation_state = 0
	}
	return nil
}

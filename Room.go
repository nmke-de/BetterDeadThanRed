package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"image/color"
)

type Room struct {
	// TODO scene data
	width   uint
	height  uint
	surface *ebiten.Image
	actors  *[]Actor // (interface for player character, mobs, NPCs, perhaps even obstacles etc)
	cache   RoomCache
	bgm     *audio.Player
}

type RoomCache map[string]int

func newRoom(w, h uint, actors *[]Actor) Room {
	return Room{
		w, h,
		ebiten.NewImage(int(w+30), int(h+30)),
		actors,
		RoomCache(map[string]int{}),
		loadOGG("glory_glory_hallelujah.ogg", audioContext),
	}
}

func (r Room) Draw(screen *ebiten.Image, _ *Game) {
	// TODO draw
	bg := color.RGBA{uint8(10), uint8(10), uint8(10), 200}
	// fg := color.RGBA{uint8(255), uint8(255), uint8(255), 200}
	screen.Fill(bg)
	r.surface.Fill(bg)
	for _, a := range *r.actors {
		a.Draw(r.surface)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(15, 0)
	screen.DrawImage(r.surface, op)
}

func (r Room) Update(game *Game, pressed []ebiten.Key) error {
	// Check whether audio is playing
	if !r.bgm.IsPlaying() {
		err := r.bgm.Rewind()
		unwrap(err)
		r.bgm.Play()
	}

	// Update cache
	for i, a := range *r.actors {
		isplayer := false
		for _, allegiance := range a.Allegiance() {
			if allegiance == player {
				isplayer = true
			}
		}
		if isplayer {
			r.cache["player"] = i
		}
	}

	dead := []int{}
	for i, a := range *r.actors {
		unwrap(a.Update())
		// Collision detection and resolve
		ax, ay := a.Position()
		for _, a2 := range *r.actors {
			a2x, a2y := a2.Position()
			if ax == a2x && ay == a2y {
				continue
			}
			dist := distance(ax, ay, a2x, a2y)
			if dist < (a.Hitbox() + a2.Hitbox()) {
				a.Collide(a2)
				a2.Collide(a)
			}
		}

		// Detect dead actors
		if !a.Alive() {
			dead = append(dead, i)
		}
	}

	// Remove dead actors
	for _, i := range dead {
		*r.actors = remove(*r.actors, i)
	}

	return nil
}

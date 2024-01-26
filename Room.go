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
	bgm := loadOGG("glory_glory_hallelujah.ogg", audioContext)
	bgm.SetVolume(0.02)
	return Room{
		w, h,
		ebiten.NewImage(int(w+30), int(h+30)),
		actors,
		RoomCache(map[string]int{}),
		bgm,
	}
}

func (r Room) Draw(screen *ebiten.Image, _ *Game) {
	// Check whether audio is playing
	if !r.bgm.IsPlaying() {
		err := r.bgm.Rewind()
		unwrap(err)
		r.bgm.Play()
	}

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
		ah := a.Hitbox()
		for _, a2 := range *r.actors {
			a2x, a2y := a2.Position()
			a2h := a2.Hitbox()
			if ax == a2x && ay == a2y {
				continue
			}
			dist := distance(ax+ah, ay+ah, a2x+a2h, a2y+a2h)
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
	removed_dead := 0
	for _, i := range dead {
		*r.actors = remove(*r.actors, i-removed_dead)
		removed_dead++
	}

	return nil
}

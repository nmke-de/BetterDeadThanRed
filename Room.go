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
		RoomCache(map[string]int{
			"spawntick":    0,
			"dead_commies": 0,
		}),
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
			px, py := a.Position()
			r.cache["px"] = int(px)
			r.cache["py"] = int(py)
		}
	}
	r.cache["spawntick"] = (r.cache["spawntick"] + 1) % 200

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
			for _, v := range a.Allegiance() {
				if v == player {
					game.current = scenes["GameOver"]
				} else if v == commie {
					r.cache["dead_commies"] += 1
				}
			}
		}
	}

	// Remove dead actors
	removed_dead := 0
	for _, i := range dead {
		*r.actors = remove(*r.actors, i-removed_dead)
		removed_dead++
	}

	// Spawn new enemies
	if r.cache["spawntick"] == 0 {
		if r.cache["dead_commies"] < 12 {
			*r.actors = append(*r.actors, newCommie((uint(r.cache["px"])+250)%r.width, (uint(r.cache["py"])+250)%r.height, "Room"))
		} else if r.cache["dead_commies"] < 17 {
			*r.actors = append(*r.actors, newComradeVodka((uint(r.cache["px"])+250)%r.width, (uint(r.cache["py"])+250)%r.height, "Room"))
		}
	}

	if r.cache["dead_commies"] == 18 {
		println("Victory")
		game.current = scenes["Victory"]
	}

	return nil
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type Room struct {
	// TODO scene data
	width  uint
	height uint
	actors *[]Actor // (interface for player character, mobs, NPCs, perhaps even obstacles etc)
	cache  RoomCache
}

type RoomCache map[string]int

func (r Room) Draw(screen *ebiten.Image, _ *Game) {
	// TODO draw
	bg := color.RGBA{uint8(10), uint8(10), uint8(10), 200}
	// fg := color.RGBA{uint8(255), uint8(255), uint8(255), 200}
	screen.Fill(bg)
	for _, a := range *r.actors {
		a.Draw(screen)
	}
}

func (r Room) Update(game *Game, pressed []ebiten.Key) error {
	// TODO update
	/*for _, key := range pressed {
		if key == ebiten.KeySpace {
			game.current = scenes[t.next_scene]
		}
	}*/
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

	for i, a := range *r.actors {
		unwrap(a.Update(r))

		// Remove dead actors
		if !a.Alive() {
			*r.actors = remove(*r.actors, i)
		}
	}
	return nil
}

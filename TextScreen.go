package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type TextScreen struct {
	title      string
	text       string
	next_scene string
}

func (t TextScreen) Draw(screen *ebiten.Image, _ *Game) {
	bg := color.RGBA{uint8(10), uint8(10), uint8(10), 200}
	fg := color.RGBA{uint8(255), uint8(255), uint8(255), 200}
	screen.Fill(bg)
	text.Draw(screen, t.title, big_fontface, 10, 50, fg)
	text.Draw(screen, t.text, fontface, 15, 100, fg)
}

func (t TextScreen) Update(game *Game, pressed []ebiten.Key) error {
	for _, key := range pressed {
		if key == ebiten.KeySpace {
			game.current = scenes[t.next_scene]
		}
	}
	return nil
}

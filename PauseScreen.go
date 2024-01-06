package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type PauseScreen struct {
	next Scene
}

func (p PauseScreen) Draw(screen *ebiten.Image, g *Game) {
	bg := color.RGBA{uint8(10), uint8(10), uint8(10), 200}
	fg := color.RGBA{uint8(255), uint8(255), uint8(255), 200}
	p.next.Draw(screen, g)
	overlay := ebiten.NewImage(layout_width, layout_height)
	overlay.Fill(bg)
	text.Draw(overlay, "PAUSED", big_fontface, 10, 50, fg)
	text.Draw(overlay, "Press Space to continue", fontface, 15, 100, fg)
	screen.DrawImage(overlay, nil)
}

func (p PauseScreen) Update(game *Game, pressed []ebiten.Key) error {
	game.paused = true
	for _, key := range pressed {
		if key == ebiten.KeySpace {
			game.current = p.next
			game.paused = false
		}
	}
	return nil
}

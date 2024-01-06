package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type FunkyHello struct {
	limit      int
	next_scene string
}

func (f FunkyHello) Draw(screen *ebiten.Image, g *Game) {
	limit := f.limit
	red_bg := uint8(g.x % 256)
	red_fg := uint8(256 - g.x%256)
	green_bg := uint8(g.x%128 + 64)
	green_fg := uint8((g.x%128 + 192) % 256)
	blue_bg := uint8(limit * g.x % 256)
	blue_fg := uint8((g.x / limit) % 256)
	screen.Fill(color.RGBA{red_bg, green_bg, blue_bg, 100})
	for i := 0; i < limit; i++ {
		text.Draw(screen, "Hello, there!", fontface, (g.x+i*(layout_width/limit))%layout_width, (i+1)*(layout_height/limit)-10, color.RGBA{red_fg, green_fg, blue_fg, 100})
	}
}

func (f FunkyHello) Update(game *Game, pressed []ebiten.Key) error {
	game.x++
	for _, key := range pressed {
		if key == ebiten.KeySpace {
			game.current = scenes[f.next_scene]
		}
	}
	return nil
}

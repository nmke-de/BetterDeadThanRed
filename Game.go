package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	current Scene
	// Needed by FunkyHello
	x int
	// Whether paused or not
	paused bool
}

func (g *Game) Update() error {
	pressed := inpututil.AppendJustPressedKeys([]ebiten.Key{})
	if !g.paused {
		for _, key := range pressed {
			if key == ebiten.KeyEscape {
				g.current = PauseScreen{g.current}
			}
		}
	}
	return g.current.Update(g, pressed)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.current.Draw(screen, g)
}

func (g *Game) Layout(width, height int) (int, int) {
	return layout_width, layout_height
}

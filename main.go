package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	layout_width  = 600
	layout_height = 480
)

var big_fontface font.Face
var fontface font.Face

func init() {
	// Font
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	unwrap(err)
	big_fontface, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	unwrap(err)
	fontface, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	unwrap(err)
}

var audioContext *audio.Context
const sampleRate = 44100

func init() {
	audioContext = audio.NewContext(sampleRate)
}

var scenes map[string]Scene

func init() {
	scenes = map[string]Scene{
		"title":   TextScreen{"La TITELO", "Hello, there!\nLorem ipsum dolor sit, amet. Consetetur.", "Funky12"},
		"Funky12": FunkyHello{12, "Funky3"},
		"Victory": TextScreen{"Victory!", "If you see this, you have won!\nCongrats!", "Room"},
		"Funky3":  FunkyHello{3, "Victory"},
		"Room": newRoom(
			550, 450,
			&[]Actor{
				newPlayer(100, 100, "Room"),
				newComradeVodka(300, 300, "Room"),
				newCommie(400, 400, "Room"),
			},
		),
	}
}

const on_frames = 7

func main() {
	ebiten.SetWindowSize(layout_width, layout_height)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Better Dead Than Red")
	game := &Game{scenes["title"], 0, false}
	err := ebiten.RunGame(game)
	unwrap(err)
}

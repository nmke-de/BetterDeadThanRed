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
		"1":        TextScreen{"", "The most innovative quadrant presents …\n\n\n\n(Press space)", "2"},
		"2":        TextScreen{"BETTER DEAD THAN RED", "", "Jan45.2"},
		"Jan45.2":  TextScreen{"January 1945, East Prussia", "", "Jan45.3"},
		"Jan45.3":  TextScreen{"January 1945, East Prussia", "Dire times have come to Europe. After the Germans have conquered\nmost of the European continent, they are now on the retreat. As\nlong as they had been around, one could have a good time if one is\nGerman and not a political dissident. Otherwise, one had had a hard\ntime.", "Jan45.4"},
		"Jan45.4":  TextScreen{"January 1945, East Prussia", "Dire times have come to Europe. After the Germans have conquered\nmost of the European continent, they are now on the retreat. As\nlong as they had been around, one could have a good time if one is\nGerman and not a political dissident. Otherwise, one had had a hard\ntime.\n\nIn stead of the Nazis, the Red Army is now on the advance. Their\nadvance is propagandized as ”liberation“. It is truly a liberation\nfrom the tyranny of life. Their leader, Josef Stalin, certainly acts\nlike a tyrant himself.", "Jan45.5"},
		"Jan45.5":  TextScreen{"January 1945, East Prussia", "Millions, whose ancestors have lived in these lands for centuries,\nhave packed their things and are fleeing westwards. Those who are\ntoo slow are killed by the Reds. Those who aren't are often killed\nfrom the coldest winter in a century. Those who turn to the Soviets\nare hanged as traitors. Those who somehow make it are not wel-\ncomed by their Western brethren.", "Jan45.6"},
		"Jan45.6":  TextScreen{"January 1945, East Prussia", "Millions, whose ancestors have lived in these lands for centuries,\nhave packed their things and are fleeing westwards. Those who are\ntoo slow are killed by the Reds. Those who aren't are often killed\nfrom the coldest winter in a century. Those who turn to the Soviets\nare hanged as traitors. Those who somehow make it are not wel-\ncomed by their Western brethren.\n\nJOHANN, a local German, has pondered his options. His father, a\nstore owner, was hanged for refusing the Hitler salute. His mother,\nwhose family had fled the Bolshevists during the revolution, has\nrecently been abducted to Auschwitz, because her origins made the\nlocal Gauleiter suspicious that she would turn to the Russians.", "Jan45.7"},
		"Jan45.7":  TextScreen{"January 1945, East Prussia", "JOHANN has thought about taking the shortcut over the partially\nfrozen Baltic Sea, but decided against that. Instead, he stands in a\ndark forest, after looting an abandoned German fortification for a\nmachine gun. With the Nazis gone, he finally has liberty, and he shall\nnot surrender his liberty to these commies!", "Tutorial"},
		"Tutorial": TextScreen{"Tutorial", "Press Escape to pause.\nPress WASD to move.\nPress Arrow Keys to shoot.\nGood luck!", "Room"},
		"Room": newRoom(
			550, 450,
			&[]Actor{
				newPlayer(100, 100, "Room"),
			},
		),
		"GameOver": TextScreen{"GAME OVER", "JOHANN died while valiantly fighting off commies.\nJoin LibRight to get better!", "GameOver"},
		"Victory":  TextScreen{"VICTORY", "JOHANN fought off the commies for now, so he sets out to find\nlike minded people, united in their strive for freedom.\nFight with JOHANN by joining LibRight!", "Victory"},
	}
}

const on_frames = 7

func main() {
	ebiten.SetWindowSize(layout_width, layout_height)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Better Dead Than Red")
	game := &Game{scenes["1"], 0, false}
	err := ebiten.RunGame(game)
	unwrap(err)
}

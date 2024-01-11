package main

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func loadOGG(path string, context *audio.Context) *audio.Player {
	f, err := assets.Open(path)
	unwrap(err)
	defer f.Close()
	stream, err := vorbis.DecodeWithoutResampling(f)
	unwrap(err)
	player, err := context.NewPlayer(stream)
	unwrap(err)
	return player
}

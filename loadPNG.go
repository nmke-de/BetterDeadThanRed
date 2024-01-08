package main

import (
	"image"
	_ "image/png"
)

func loadPNG(path string) image.Image {
	f, err := assets.Open(path)
	unwrap(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	unwrap(err)
	return img
}

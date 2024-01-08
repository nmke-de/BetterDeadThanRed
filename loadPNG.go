package main

import (
	"os"
	"image"
	_ "image/png"
)

func loadPNG(path string) image.Image {
	f, err := os.Open(path)
	unwrap(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	unwrap(err)
	return img
}

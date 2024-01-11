package main

import "embed"

//go:embed *.png
//go:embed *.ogg
var assets embed.FS

package main

import "math"

func distance(x1, y1, x2, y2 uint) uint {
	return uint(math.Floor(math.Sqrt(float64(
		(int(x2)-int(x1))*(int(x2)-int(x1)) +
			(int(y2)-int(y1))*(int(y2)-int(y1)),
	))))
}

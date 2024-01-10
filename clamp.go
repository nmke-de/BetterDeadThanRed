package main

func clamp(val, minimum, maximum uint) uint {
	return uint(min(max(int(val), int(minimum)), int(maximum)))
}

package main

func remove[T any](arr []T, index int) []T {
	if len(arr) == 0 {
		return arr
	}
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

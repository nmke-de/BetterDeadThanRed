package main

func remove(arr []any, index uint) []any {
	if len(arr) == 0 {
		return arr
	}
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

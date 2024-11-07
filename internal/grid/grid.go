package grid

import "math/rand"

// CreateGrid creates a 2d slice of runes with a default rune in each element.
func CreateGrid(size int, def rune) [][]rune {
	arr := make([][]rune, size)
	for i := range arr {
		arr[i] = make([]rune, size)
		for j := range arr[i] {
			arr[i][j] = def
		}
	}
	return arr
}

// ReplaceCharacters replaces any "from" runes to a random letter in the whole grid
func ReplaceCharacters(grid [][]rune, from rune) {
	runes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == from {
				grid[i][j] = runes[rand.Intn(len(runes))]
			}
		}
	}
}

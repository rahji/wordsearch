package wordsearch

import (
	"math/rand"
)

const (
	defaultRune = '.'
	alphabet    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// direction is a private type that represents a cardinal direction
type direction struct {
	x int
	y int
}

// allDirections is a private helper function that returns a slice
// of all legitimate `direction` values
func allDirections() []direction {
	return []direction{
		{x: 0, y: 1},   // N
		{x: 1, y: 1},   // NE
		{x: 1, y: 0},   // E
		{x: 1, y: -1},  // SE
		{x: 0, y: -1},  // S
		{x: -1, y: -1}, // SW
		{x: -1, y: 0},  // W
		{x: -1, y: 1},  // NW
	}
}

// WordSearch contains config info and the actual grid of rows and cols
type WordSearch struct {
	Size       int
	directions []direction
	Grid       [][]rune
}

// createEmptyGrid creates a 2d slice of runes with the default rune in each element.
func createEmptyGrid(size int) [][]rune {
	arr := make([][]rune, size)
	for i := range arr {
		arr[i] = make([]rune, size)
		for j := range arr[i] {
			arr[i][j] = defaultRune
		}
	}
	return arr
}

// NewWordSearch initializes and returns a WordSearch instance
func NewWordSearch(size int) *WordSearch {
	return &WordSearch{
		Size:       size,
		directions: allDirections(),
		Grid:       createEmptyGrid(size),
	}
}

// fillGrid replaces any default runes to a random letter in the whole grid
func (ws *WordSearch) fillGrid() {
	letters := []rune(alphabet)
	for i := range ws.Grid {
		for j := range ws.Grid[i] {
			if ws.Grid[i][j] == defaultRune {
				ws.Grid[i][j] = letters[rand.Intn(len(letters))]
			}
		}
	}
}

/*
// CanPlaceWord returns true if the word can be placed in the specified location on the grid
func CanPlaceWord(grid [][]rune, word []rune, row int, col int, dir Direction) (bool) {
	size := len(grid)
	for i := 0; i < len(word) ; i++ {
		r := row + i * dir.y
		c := row + i * dir.x
		if r < 0 || r >= size
			return false
		if c <
		    fmt.Printf("Rune %v is '%c'\n", i, runes[i])
}
}
// PlaceWord tries to place the specified word can be placed on the grid
func PlaceWord(grid [][]rune, word []rune, row int, col int, dir Direction, size int) (error) {
	for r := range grid {
		for c := range grid[r] {
			newRow := row +
		}

	}
}
*/

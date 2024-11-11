package wordsearch

import (
	"errors"
	"math/rand"
)

const (
	defaultRune = '.'
	alphabet    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// direction is a private type that represents the 2 axes of a cardinal direction
type direction struct {
	x int
	y int
}

// allDirectionsMap returns a map where keys are string representations
// of cardinal directions and values are `direction` structs
func allDirectionsMap() map[string]direction {
	return map[string]direction{
		"N":  {x: 0, y: -1},
		"NE": {x: 1, y: -1},
		"E":  {x: 1, y: 0},
		"SE": {x: 1, y: 1},
		"S":  {x: 0, y: 1},
		"SW": {x: -1, y: 1},
		"W":  {x: -1, y: 0},
		"NW": {x: -1, y: -1},
	}
}

// WordSearch contains config info and the actual grid of rows and cols
type WordSearch struct {
	Size       int
	directions map[string]direction
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
		directions: allDirectionsMap(),
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

// PlaceWord tries to write a word to the grid and returns an error if it isn't a valid location
func (ws *WordSearch) PlaceWord(str string, row int, col int, cardinal string) error {
	dir := ws.directions[cardinal]
	word := []rune(str)
	tempGrid := ws.Grid
	for i := 0; i < len(word); i++ {
		r := row + i*dir.y
		c := col + i*dir.x
		if r < 0 || r >= ws.Size || c < 0 || c >= ws.Size {
			return errors.New("word extends outside of the grid")
		}
		if ws.Grid[r][c] != defaultRune {
			return errors.New("word would overlap an existing word")
		}
		tempGrid[r][c] = word[i]
	}
	ws.Grid = tempGrid
	return nil
}

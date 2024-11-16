package wordsearch

import (
	"errors"
	"math/rand"
	"sort"
)

const (
	defaultRune = '.'
	alphabet    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	attempts    = 100
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

// filterDirectionsMap returns a subset of the map that allDirectionsMap returns
func filterDirectionsMap(original map[string]direction, filter []string) map[string]direction {
	filtered := make(map[string]direction)
	for _, key := range filter {
		if val, exists := original[key]; exists {
			filtered[key] = val
		}
	}
	return filtered
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
func NewWordSearch(size int, cardinals []string) *WordSearch {
	dirs := make(map[string]direction)
	if cardinals == nil {
		dirs = allDirectionsMap()
	} else {
		dirs = filterDirectionsMap(allDirectionsMap(), cardinals)
	}
	return &WordSearch{
		Size:       size,
		directions: dirs,
		Grid:       createEmptyGrid(size),
	}
}

// PlaceWord tries to write a word to a specific place on the grid in a specific direction.
// It returns an error if it can't be done for some reason.
// It is assumed that Placeword is called on longer words before being called on shorter words.
// This should keep a shorter word from being completely inside a longer one
func (ws *WordSearch) PlaceWord(str string, row int, col int, cardinal string) error {
	dir := ws.directions[cardinal]
	word := []rune(str)
	tempGrid := ws.Grid // make a temp grid so a failed attempt doesn't corrupt the real one
	overlapCount := 0   // the number of valid overlapping letters (a complete overlap of words is invalid)
	// loop through each rune of the word
	for i := 0; i < len(word); i++ {
		r := row + i*dir.y
		c := col + i*dir.x
		if r < 0 || r >= ws.Size || c < 0 || c >= ws.Size {
			return errors.New("word extends outside of the grid")
		}
		if ws.Grid[r][c] != defaultRune && ws.Grid[r][c] != word[i] {
			return errors.New("a letter would overwrite an existing (different) letter")
		}
		if ws.Grid[r][c] == word[i] {
			overlapCount++
		}
		if overlapCount == len(word) {
			return errors.New("word would be completely inside another word")
		}
		tempGrid[r][c] = word[i]
	}
	ws.Grid = tempGrid
	return nil
}

// fillRemainingGrid replaces any default runes to a random letter in the whole grid
func (ws *WordSearch) fillRemainingGrid() {
	letters := []rune(alphabet)
	for i := range ws.Grid {
		for j := range ws.Grid[i] {
			if ws.Grid[i][j] == defaultRune {
				ws.Grid[i][j] = letters[rand.Intn(len(letters))]
			}
		}
	}
}

// CreatePuzzle places the words from the words list and returns a list
// of words that could not be placed
func (ws *WordSearch) CreatePuzzle(words []string) (unplaced []string) {
	// sort the slice of words by length, longest first
	// this is to avoid a shorter word being placed entirely within another longer word
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	// make a slice of strings containing the keys to the direction map (eg: "N", "SE", etc.)
	keys := make([]string, 0, len(ws.directions))
	for key := range ws.directions {
		keys = append(keys, key)
	}
	// make a bunch of random attempts to fit each word into the grid
	for _, word := range words {
		placed := false
		for i := 0; i < attempts; i++ {
			dir := keys[rand.Intn(len(keys))] // `dir` is "N", "SE", etc.
			row := rand.Intn(ws.Size - 1)
			col := rand.Intn(ws.Size - 1)
			err := ws.PlaceWord(word, row, col, dir)
			if err == nil {
				placed = true
				break
			}
		}
		if placed == false {
			unplaced = append(unplaced, word)
		}
	}
	ws.fillRemainingGrid()
	return
}

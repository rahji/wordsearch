package wordsearch

import (
	"errors"
	"math/rand"
	"sort"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	attempts = 100
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

// filterDirectionsMap returns a subset of the map that allDirectionsMap returns.
// This is used to create a limited map of `direction` eg: to eliminate diagonal placement, etc.
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
	Grid       [][]byte
}

// isUppercase returns true if the byte is a capital letter.
// Capital letters are intentionally placed letters, while lowercase letters are randomly placed and
// can be overwritten with intentionally placed letters. Everything in the grid should either be a
// lowercase or uppercase letter.
func isUppercase(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// toLowercase turns an uppercase byte into lowercase
func toLowercase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

// createEmptyGrid creates a 2d slice of bytes with random lowercase letters in each element.
// Lowercase letters represent letters that were not placed intentionally.
func createEmptyGrid(size int) [][]byte {
	arr := make([][]byte, size)
	for i := range arr {
		arr[i] = make([]byte, size)
		for j := range arr[i] {
			randomIndex := rand.Intn(len(alphabet))
			arr[i][j] = toLowercase(alphabet[randomIndex])
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
func (ws *WordSearch) PlaceWord(word string, row int, col int, cardinal string) error {
	dir := ws.directions[cardinal]
	tempGrid := ws.Grid // make a temp grid so a failed attempt doesn't corrupt the real one
	overlapCount := 0   // the number of valid overlapping letters (a complete overlap of words is invalid)
	// loop through each byte of the word
	for i := 0; i < len(word); i++ {
		r := row + i*dir.y
		c := col + i*dir.x
		if r < 0 || r >= ws.Size || c < 0 || c >= ws.Size {
			return errors.New("word extends outside of the grid")
		}
		if isUppercase(ws.Grid[r][c]) && ws.Grid[r][c] != word[i] {
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
	return
}

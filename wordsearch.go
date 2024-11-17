// This package is the basis for a word search puzzle generator
//
// It generates a grid of letters containing a list of hidden words.
// Configuration includes an optional list of cardinal directions (e.g. "N", "SW", etc.)
// in which words can be placed, the size of the grid, and whether
// overlapping letters are allowed. The grid is a 2D slice of bytes
// containing lowercase letters for "filler" letters and
// uppercase letters for words that have been explicitly placed.
// A helper function can return the grid in other formats.
package wordsearch

import (
	"errors"
	"math/rand"
	"sort"
	"strings"

	"github.com/rahji/wordsearch/internal/letters"
	"github.com/rahji/wordsearch/internal/vector"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	attempts = 100 // max number of times to attempt to place a word
)

// Gridstyle is an enum-like list of ways that the output of ReturnGrid can be styled
type GridStyle int

const (
	GridRaw GridStyle = iota
	GridWithDots
	GridWithSpaces
	GridAllUppercase
	GridAllLowercase
)

// WordSearch is a struct that contains the puzzle configuration and the actual grid of rows and cols,
// which can be accessed directly or via the helper method ReturnGrid. The config includes the
// size of the puzzle, allowable directions (as one- or two-letter abbreviations for the cardinal directions),
// and whether overlapping is allowed
type WordSearch struct {
	Size       int
	Grid       [][]byte
	Directions []string
	Overlaps   bool
}

// createEmptyGrid creates a 2d slice of bytes with random lowercase letters in each element.
// Lowercase letters represent letters that were not placed intentionally.
func createEmptyGrid(size int) [][]byte {
	arr := make([][]byte, size)
	for i := range arr {
		arr[i] = make([]byte, size)
		for j := range arr[i] {
			randomIndex := rand.Intn(len(alphabet))
			arr[i][j] = letters.ToLowercase(alphabet[randomIndex])
		}
	}
	return arr
}

// NewWordSearch initializes and returns a WordSearch instance.
// The size parameter is both the width and height of the grid.
// The cardinals parameter is either a slice of strings that are
// abbreviations for the cardinal directions (N, NE, E, SE, S, SW, W, NW)
// or nil for all possible directions.
// The overlaps parameter determines whether any overlapping of words is allowed.
func NewWordSearch(size int, cardinals []string, overlaps bool) *WordSearch {
	if cardinals == nil {
		cardinals = []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	}
	return &WordSearch{
		Size:       size,
		Directions: cardinals,
		Overlaps:   overlaps,
		Grid:       createEmptyGrid(size),
	}
}

// ReturnGrid returns the grid, with the bytes restyled using a parameter of the GridStyle type
func (ws *WordSearch) ReturnGrid(style GridStyle) [][]byte {
	if style == GridRaw {
		return ws.Grid
	}

	returnGrid := make([][]byte, len(ws.Grid))
	for r := range ws.Grid {
		returnGrid[r] = make([]byte, len(ws.Grid[r]))
	}

	// if the lowercase letters are going to be replaced with symbols...
	replacementChar := byte(0)
	switch style {
	case GridWithDots:
		replacementChar = '.'
	case GridWithSpaces:
		replacementChar = ' '
	}

	// loop through the grid and replace either lowercase or uppercase letters...
	for i, row := range ws.Grid {
		for j, b := range row {
			returnGrid[i][j] = b
			// replace lowercase byte with a symbol if that's the style
			if replacementChar != 0 && letters.IsLowercase(b) {
				returnGrid[i][j] = replacementChar
			}
			if style == GridAllLowercase {
				returnGrid[i][j] = letters.ToLowercase(b)
			}
			if style == GridAllUppercase {
				returnGrid[i][j] = letters.ToUppercase(b)
			}
		}
	}
	return returnGrid
}

// PlaceWord tries to write a single word to a specific place on the grid in a specific direction.
// This function is where the word gets capitalized. It is assumed to be a word made only of the letters A-Z.
// It returns an error if it can't be done for some reason. The possible reasons for failure are:
//  1. The placement would extend outside of the grid
//  2. A letter in the word would overwrite an existing (different) letter
//  3. A letter overlaps another placed letter and overlaps are disallowed in this word search
//  4. Overlaps are alllowed, but the word would be placed completely inside another word (which is never allowed)
func (ws *WordSearch) PlaceWord(word string, row int, col int, cardinal string) error {
	dir := vector.CardinalToVector(cardinal)
	overlapCount := 0 // the number of valid overlapping letters (a complete overlap of words is invalid)

	// make a deep copy of the grid so a failed attempt won't screw up the real grid
	tempGrid := make([][]byte, len(ws.Grid))
	for i := range ws.Grid {
		tempGrid[i] = make([]byte, len(ws.Grid[i]))
		copy(tempGrid[i], ws.Grid[i])
	}
	// loop through each byte of the word
	word = strings.ToUpper(word)
	for i := 0; i < len(word); i++ {
		r := row + i*dir.Y
		c := col + i*dir.X
		if r < 0 || r >= ws.Size || c < 0 || c >= ws.Size {
			return errors.New("word extends outside of the grid")
		}
		if letters.IsUppercase(ws.Grid[r][c]) && ws.Overlaps == false {
			return errors.New("a letter would overlap another letter and overlaps are disallowed")
		}
		if letters.IsUppercase(ws.Grid[r][c]) && ws.Grid[r][c] != word[i] {
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

// CreatePuzzle places words from a words list, after sorting them by length, longest first.
// It returns nil if successful. Otherwise it returns a slice of words that could not be placed
// after the maximum number of attempts.
func (ws *WordSearch) CreatePuzzle(words []string) (unplaced []string) {
	// sort the slice of words by length, longest first
	// this is to avoid a shorter word being placed entirely within another longer word
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	// make a bunch of random attempts to fit each word into the grid
	for _, word := range words {
		placed := false
		for i := 0; i < attempts; i++ {
			randomIndex := rand.Intn(len(ws.Directions))
			randomCardinal := ws.Directions[randomIndex]
			row := rand.Intn(ws.Size - 1)
			col := rand.Intn(ws.Size - 1)
			err := ws.PlaceWord(word, row, col, randomCardinal)
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

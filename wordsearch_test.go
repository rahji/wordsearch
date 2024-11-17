package wordsearch

import (
	"testing"
)

// printGrid is a private function that logs the grid
func printGrid(t *testing.T, grid [][]byte) {
	for _, cell := range grid {
		t.Log(string(cell))
	}
}

// TestCreateEmptyGrid creates a WordSearch instance and verifies that an empty grid was created
func TestCreateEmptyGrid(t *testing.T) {

	ws := NewWordSearch(15, nil, true)

	t.Run("Check if the row slice has the correct length", func(t *testing.T) {
		if len(ws.Grid) != 15 {
			t.Errorf("Expected grid length of 15, got %d", len(ws.Grid))
		}
	})

	t.Run("Check if the column slice has the correct length", func(t *testing.T) {
		for i, row := range ws.Grid {
			if len(row) != 15 {
				t.Errorf("Row %d: Expected length of 15, got %d", i, len(row))
			}
		}
	})

	t.Run("Check if all elements are lowercase letters", func(t *testing.T) {
		for i, row := range ws.Grid {
			for j, char := range row {
				if !(char >= 'a' && char <= 'z') {
					t.Errorf("Position [%d][%d]: Expected lowercase letter, got %c", i, j, char)
				}
			}
		}
	})
}

// TestPlaceWord tests the placement of the word FOUR in various positions and directions
// tests: horizontal, vertical, too far in each of the cardinal directions, overlap!
func TestPlaceWord(t *testing.T) {
	tests := []struct {
		name      string
		row, col  int
		direction string
		wantError bool
	}{
		{
			name:      "Place the word FOUR horizontally",
			row:       0,
			col:       0,
			direction: "E",
			wantError: false,
		},
		{
			name:      "Place the word FOUR vertically",
			row:       9,
			col:       0,
			direction: "N",
			wantError: false,
		},
		{
			name:      "Place the word FOUR exceeding the left boundary",
			row:       0,
			col:       2,
			direction: "W",
			wantError: true,
		},
		{
			name:      "Place the word FOUR exceeding the right boundary",
			row:       0,
			col:       8,
			direction: "E",
			wantError: true,
		},
		{
			name:      "Place the word FOUR exceeding the top boundary",
			row:       0,
			col:       2,
			direction: "N",
			wantError: true,
		},
		{
			name:      "Place the word FOUR exceeding the bottom boundary",
			row:       8,
			col:       2,
			direction: "S",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := NewWordSearch(10, nil, true)
			err := ws.PlaceWord("FOUR", tt.row, tt.col, tt.direction)
			if (err != nil) != tt.wantError {
				t.Errorf("PlaceWord() error = %v, wantError %v got %v", err, tt.wantError, err != nil)
			}
			t.Log(tt.name)
			printGrid(t, ws.Grid)
		})
	}
}

// TestOverlappingWords tests specifically for different word overlap cases
func TestOverlappingWords(t *testing.T) {

	// Make a wordsearch struct that temporarily allows overlaps.
	// The Overlaps field will be changed directly depending on the test below
	// We're creating the ws variable here because each test builds on the previous one
	// so we don't want to recreate ws with each test!
	ws := NewWordSearch(10, nil, true)

	tests := []struct {
		name      string
		word      string
		row       int
		col       int
		dir       string
		overlap   bool
		wantError bool
	}{
		{
			name:      "Place starting word FOODIE, no overlaps yet",
			word:      "FOODIE",
			row:       0,
			col:       0,
			dir:       "S",
			overlap:   true,
			wantError: false,
		},
		{
			name:      "Place OOF with single letter overlap onto FOODIE",
			word:      "OOF",
			row:       1,
			col:       0,
			dir:       "E",
			overlap:   true,
			wantError: false,
		},
		{
			name:      "Place DIES overlapping onto the end of FOODIE",
			word:      "DIES",
			row:       3,
			col:       0,
			dir:       "S",
			overlap:   true,
			wantError: false,
		},
		{
			name:      "Try to wrongly place DIE within DIES",
			word:      "DIE",
			row:       3,
			col:       0,
			dir:       "S",
			overlap:   true,
			wantError: true,
		},
		{
			name:      "Try to overlap the F in FOO onto FOODIE when overlaps are disallowed",
			word:      "FOO",
			row:       0,
			col:       0,
			dir:       "E",
			overlap:   false,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws.Overlaps = tt.overlap
			err := ws.PlaceWord(tt.word, tt.row, tt.col, tt.dir)
			if (err != nil) != tt.wantError {
				t.Errorf("Overlapping overlaps = %v, error = %v, wantError %v got %v", ws.Overlaps, err, tt.wantError, err != nil)
			}
			t.Log(tt.name)
			printGrid(t, ws.Grid)
		})
	}

}

// TestCreatePuzzle tests for proper filling out of a complete puzzle with a list of words
func TestCreatePuzzle(t *testing.T) {

	tests := []struct {
		name           string
		wordsearch     WordSearch
		words          []string
		expectUnplaced bool
	}{
		{
			name:           "normal 8x8 grid: ONE TWO THREE FOUR",
			wordsearch:     *NewWordSearch(8, nil, true),
			words:          []string{"ONE", "TWO", "THREE", "FOUR"},
			expectUnplaced: false,
		},
		{
			name:           "impossible 3x3 grid: ONE OOO TWO DOS PRO",
			wordsearch:     *NewWordSearch(3, nil, true),
			words:          []string{"ONE", "OOO", "TWO", "DOS", "PRO"},
			expectUnplaced: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unplaced := tt.wordsearch.CreatePuzzle(tt.words)
			if len(unplaced) > 0 && !tt.expectUnplaced {
				t.Errorf("expected no unplaced, got %v", len(unplaced))
			}
			// fmt.Printf("%d unplaced: %v\n", len(unplaced), unplaced)
			printGrid(t, tt.wordsearch.Grid)
		})
	}
}

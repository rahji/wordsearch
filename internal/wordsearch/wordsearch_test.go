package wordsearch

import (
	"strings"
	"testing"
)

// TestCreateEmptyGrid creates a WordSearch instance and verifies that an empty grid was created
func TestCreateEmptyGrid(t *testing.T) {

	ws := NewWordSearch(15)

	// test 1: Check if the outer slice has correct length
	if len(ws.Grid) != 15 {
		t.Errorf("Expected grid length of 15, got %d", len(ws.Grid))
	}

	// test 2: Check if each inner slice has correct length
	for i, row := range ws.Grid {
		if len(row) != 15 {
			t.Errorf("Row %d: Expected length of 15, got %d", i, len(row))
		}
	}

	// test 3: Check if all elements are asterisks
	for i, row := range ws.Grid {
		for j, char := range row {
			if char != '.' {
				t.Errorf("Position [%d][%d]: Expected '*', got %c", i, j, char)
			}
		}
	}
}

// TestFillGridVariousSizes tests that different sized grids can be filled with letters.
// It also shows that multiple WordSearch instances can be created.
func TestFillGridVariousSizes(t *testing.T) {

	tests := []struct {
		name       string
		wordsearch WordSearch
	}{
		{"small grid", *NewWordSearch(5)},
		{"large grid", *NewWordSearch(24)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, cell := range tt.wordsearch.Grid {
				t.Log(string(cell))
			}

			tt.wordsearch.fillGrid()

			for _, cell := range tt.wordsearch.Grid {
				t.Log(string(cell))
			}

			// Verify all positions contain valid letters
			validChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			for i := range tt.wordsearch.Grid {
				for j := range tt.wordsearch.Grid[i] {
					index := strings.IndexRune(validChars, tt.wordsearch.Grid[i][j])
					if index == -1 {
						t.Errorf("Position [%d][%d]: Expected a letter A-Z, got %c", i, j, tt.wordsearch.Grid[i][j])
					}
				}
			}
		})
	}
}

// TestPlaceWord tests the placement of the word FOUR in various positions and directions
func TestPlaceWord(t *testing.T) {
	tests := []struct {
		name      string
		row, col  int
		direction string
		wantError bool
	}{
		{
			name:      "horizontal placement success",
			row:       0,
			col:       0,
			direction: "E",
			wantError: false,
		},
		{
			name:      "vertical placement success",
			row:       9,
			col:       0,
			direction: "N",
			wantError: false,
		},
		{
			name:      "placement fails due to boundary",
			row:       0,
			col:       2,
			direction: "W",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := NewWordSearch(10)
			err := ws.PlaceWord("FOUR", tt.row, tt.col, tt.direction)

			if (err != nil) != tt.wantError {
				t.Errorf("PlaceWord() error = %v, wantError %v", err != nil, tt.wantError)
			}

			t.Log(tt.name)
			for _, cell := range ws.Grid {
				t.Log(string(cell))
			}
		})
	}
}

// test 1: Verify that a word can be placed horizontally
// test 2: Verify that a word can be placed vertically
// test 3: Verify that a word can NOT be placed too far to the left or right
// test 4: Verify that a word can NOT be placed too far to the top or bottom
// test 5: Verify that a word can NOT be placed on top of another word

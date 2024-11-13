package wordsearch

import (
	"strings"
	"testing"
)

// printGrid is a private function that logs the grid
func printGrid(t *testing.T, grid [][]rune) {
	for _, cell := range grid {
		t.Log(string(cell))
	}
}

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

			printGrid(t, tt.wordsearch.Grid)

			tt.wordsearch.fillRemainingGrid()

			printGrid(t, tt.wordsearch.Grid)

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
// tests: horizontal, vertical, too far in each of the cardinal directions, overlap!
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
			name:      "placement fails due to left boundary",
			row:       0,
			col:       2,
			direction: "W",
			wantError: true,
		},
		{
			name:      "placement fails due to right boundary",
			row:       0,
			col:       8,
			direction: "E",
			wantError: true,
		},
		{
			name:      "placement fails due to top boundary",
			row:       0,
			col:       2,
			direction: "N",
			wantError: true,
		},
		{
			name:      "placement fails due to bottom boundary",
			row:       8,
			col:       2,
			direction: "S",
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
			printGrid(t, ws.Grid)
		})
	}
}

func TestCreatePuzzle(t *testing.T) {

	tests := []struct {
		name           string
		wordsearch     WordSearch
		words          []string
		expectUnplaced bool
	}{
		{
			name:           "normal grid",
			wordsearch:     *NewWordSearch(10),
			words:          []string{"ONE", "TWO", "THREE", "FOUR"},
			expectUnplaced: false,
		},
		{
			name:           "overlap grid",
			wordsearch:     *NewWordSearch(3),
			words:          []string{"OOO", "OOO", "OOO"},
			expectUnplaced: false,
		},
		{
			name:           "impossible grid",
			wordsearch:     *NewWordSearch(3),
			words:          []string{"ONE", "TWO", "DOS", "POO", "PRO"},
			expectUnplaced: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unplaced := tt.wordsearch.CreatePuzzle(tt.words)
			if len(unplaced) > 0 && !tt.expectUnplaced {
				t.Errorf("expected no unplaced, got %v", len(unplaced))
			}
			printGrid(t, tt.wordsearch.Grid)
		})
	}
}

package wordsearch

import (
	"strings"
	"testing"
)

func TestCreateGrid(t *testing.T) {

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

func TestReplaceCharactersVariousSizes(t *testing.T) {

	tests := []struct {
		name       string
		wordsearch WordSearch
	}{
		{"small grid", *NewWordSearch(5)},
		{"large grid", *NewWordSearch(24)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, row := range tt.wordsearch.Grid {
				t.Log(string(row))
			}

			tt.wordsearch.fillGrid()

			for _, row := range tt.wordsearch.Grid {
				t.Log(string(row))
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

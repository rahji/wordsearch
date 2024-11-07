package grid

import (
	"strings"
	"testing"
)

func TestCreateGrid(t *testing.T) {
	// call the function
	grid := CreateGrid(15, '*')

	// test 1: Check if the outer slice has correct length
	if len(grid) != 15 {
		t.Errorf("Expected grid length of 15, got %d", len(grid))
	}

	// test 2: Check if each inner slice has correct length
	for i, row := range grid {
		if len(row) != 15 {
			t.Errorf("Row %d: Expected length of 15, got %d", i, len(row))
		}
	}

	// test 3: Check if all elements are asterisks
	for i, row := range grid {
		for j, char := range row {
			if char != '*' {
				t.Errorf("Position [%d][%d]: Expected '*', got %c", i, j, char)
			}
		}
	}
}

func TestReplaceCharactersVariousSizes(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"small grid", 5},
		{"large grid", 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := make([][]rune, tt.size)
			grid = CreateGrid(tt.size, '.')

			for _, row := range grid {
				t.Log(string(row))
			}

			ReplaceCharacters(grid, '.')

			for _, row := range grid {
				t.Log(string(row))
			}

			// Verify all positions contain valid letters
			validChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			for i := range grid {
				for j := range grid[i] {
					index := strings.IndexRune(validChars, grid[i][j])
					if index == -1 {
						t.Errorf("Position [%d][%d]: Expected a letter A-Z, got %c", i, j, grid[i][j])
					}
				}
			}
		})
	}
}

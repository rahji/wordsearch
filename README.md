# Word Search

This is a Go package that can be used for making a Word Search puzzle generator or interactive game.

Documentation should be here: https://pkg.go.dev/github.com/rahji/wordsearch

## Installation

```bash
go get github.com/rahji/wordsearch@latest
```

## Usage

```go
	ws := wordsearch.NewWordSearch(16, []string{"N","E"}, true)
	unplaced := ws.CreatePuzzle(words)
	if unplaced != nil {
		fmt.Printf("These words could not be placed: %v", unplaced)
	}

	uppercaseGrid := ws.ReturnGrid(wordsearch.GridAllUppercase)
	for i := 0; i < len(uppercaseGrid); i++ {
		for j := 0; j < len(uppercaseGrid[i]); j++ {
			fmt.Printf("%c ", uppercaseGrid[i][j])
		}
		fmt.Println()
	}
```

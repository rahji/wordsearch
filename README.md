# Word Search

This is a Go package that can be used for making a Word Search puzzle generator or interactive game.

Documentation should be here: https://pkg.go.dev/github.com/rahji/wordsearch

> Note that version 2 changes the way the constructor works.

## Installation

```bash
go get github.com/rahji/wordsearch@latest
```

## Usage

```go
  words := []string{"this","that","other"}

	ws := wordsearch.NewWordSearch(16)

	unplaced := ws.CreatePuzzle(words)
	if unplaced != nil {
		fmt.Printf("These words could not be placed: %v", unplaced)
	}

	uppercaseGrid := ws.ReturnGrid(wordsearch.GridAllUppercase)
 	for i := range uppercaseGrid {
		for j := range uppercaseGrid[i] {
			fmt.Printf("%c ", uppercaseGrid[i][j])
		}
		fmt.Println()
	}
```

This example shows how options can be used to create a kid-friendly puzzle:

```go
ws := wordsearch.NewWordSearch(16,
	wordsearch.WithDirections([]string{"N","E","W","S"}),
	wordsearch.WithoutOverlap()
)
```

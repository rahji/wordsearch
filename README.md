# Word Search

This is a Go package that can be used for making a Word Search puzzle generator or interactive game.

Documentation should be here: https://pkg.go.dev/github.com/rahji/wordsearch

## Installation

```bash
go get github.com/rahji/wordsearch
```

## Usage

```go
ws := NewWordSearch(16, []string{"N","E"}, true)
unplaced := ws.CreatePuzzle([]string{"APPLES","ORANGES","CARROTS"})
if unplaced != nil {
  fmt.Printf("These words could not be placed: %v", unplaced)
}
fmt.Printf("%v", ws.ReturnGrid(GridAllUppercase))
```

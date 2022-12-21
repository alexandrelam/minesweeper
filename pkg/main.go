package main

import (
	"fmt"
	"time"

	"github.com/alexandrelam/minesweeper/pkg/game"
)

func main() {

	// start time
	start := time.Now()

	board := game.NewBoard(16, 16, 40)
	board.Display()
	board.Play(0, 0)
	board.Play(15, 0)
	board.Play(2, 2)
	board.Flag(9, 2)
	board.Flag(13, 2)
	board.Play(12, 3)
	board.Play(13, 2)
	board.Play(14, 2)
	board.Play(2, 6)
	board.Play(2, 6)
	board.Flag(1, 6)
	board.Play(2, 6)
	board.Play(1, 5)
	board.Play(1, 4)
	board.Flag(1, 3)
	board.Play(1, 4)
	board.Play(12, 2)
	board.Flag(15, 2)
	board.Play(14, 2)
	board.Flag(10, 3)
	board.Play(10, 2)

	elapsed := time.Since(start)
	fmt.Printf("ran in %d ms", elapsed.Milliseconds())
}

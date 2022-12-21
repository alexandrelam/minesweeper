package main

import (
	"github.com/alexandrelam/minesweeper/pkg/game"
)

func main() {
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
	board.Play(15, 2)
	board.DisplayNoHidden()
}

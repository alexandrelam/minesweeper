package main

import (
	"github.com/alexandrelam/minesweeper/pkg/game"
)

func main() {
	board := game.NewBoard(16, 16, 40)
	board.Display()
	board.Play(0, 0)
	board.Display()
	board.Play(15, 0)
	board.Display()
	board.Play(2, 2)
	board.Display()
}

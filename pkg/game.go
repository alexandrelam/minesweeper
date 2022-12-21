package main

import (
	"encoding/json"

	"github.com/alexandrelam/minesweeper/pkg/game"
)

func (c *Client) createGame() {
	board := game.NewBoard

	jsonSquares, err := json.Marshal(board.GetSquare())

	if err != nil {
		panic(err)
	}

	c.send <- jsonSquares
}

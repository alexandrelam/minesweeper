package main

import (
	"encoding/json"

	"github.com/alexandrelam/minesweeper/pkg/game"
)

var board *game.Board

func (c *Client) createGame() {
	board = game.NewBoard(16, 16, 40)
	c.sendBoard()
}

func (c *Client) flag(row, column int) {
	played := board.Flag(row, column)
	if played {
		c.sendBoard()
	}
}

func (c *Client) unflag(row, column int) {
	played := board.Unflag(row, column)
	if played {
		c.sendBoard()
	}
}

func (c *Client) dig(row, column int) {
	playStatus := board.Play(row, column)
	if playStatus.IsPlayed == true {
		if playStatus.IsLost == true {
			c.hub.broadcast <- []byte("BOOM!")
			return
		} else {
			c.sendBoard()
		}
	}
}

func (c *Client) sendBoard() {
	jsonSquares, err := json.Marshal(board.GetSquare())

	if err != nil {
		panic(err)
	}

	c.hub.broadcast <- jsonSquares
}

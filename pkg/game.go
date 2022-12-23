package main

import (
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
			response := newReponse(GAME_LOST, board.GetSquare())
			c.hub.broadcast <- response.toJSON()
			return
		} else {
			c.sendBoard()
		}
	}
}

func (c *Client) sendBoard() {
	response := newReponse(UPDATE_BOARD, board.GetSquare())
	c.hub.broadcast <- response.toJSON()
}

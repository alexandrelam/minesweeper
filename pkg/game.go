package main

import (
	"github.com/alexandrelam/minesweeper/pkg/game"
)

var board *game.Board

func (c *Client) createGame() {
	board = game.NewBoard(16, 16, 40)
	resetHistory()

	c.sendBoard(NilEvent)
}

func (c *Client) flag(row, column int) {
	played := board.Flag(row, column)
	event := newEvent(FLAG, c.id, c.name, row, column)

	if played {
		c.sendBoard(event)
		c.sendHistory()
	}
}

func (c *Client) unflag(row, column int) {
	played := board.Unflag(row, column)
	event := newEvent(UNFLAG, c.id, c.name, row, column)

	if played {
		c.sendBoard(event)
		c.sendHistory()
	}
}

func (c *Client) dig(row, column int) {
	playStatus := board.Play(row, column)
	event := newEvent(DIG, c.id, c.name, row, column)
	if playStatus.IsPlayed == true {
		if playStatus.IsLost == true {
			event.Action = GAME_LOST
			appendEvent(event)
			c.sendHistory()
			response := newReponse(GAME_LOST, board.GetSquare())
			c.hub.broadcast <- response.toJSON()
		} else if playStatus.IsWin {
			event.Action = GAME_WON
			appendEvent(event)
			c.sendHistory()
			response := newReponse(GAME_WON, board.GetSquare())
			c.hub.broadcast <- response.toJSON()
		} else {
			c.sendBoard(event)
		}
	}
}

func (c *Client) sendBoard(event Event) {
	response := newReponse(UPDATE_BOARD, board.GetSquare())
	c.hub.broadcast <- response.toJSON()

	if event != NilEvent {
		appendEvent(event)
		c.sendHistory()
	}
}

package main

import (
	"encoding/json"
)

const (
	CREATE_GAME = "CREATE_GAME"
	FLAG        = "FLAG"
	UNFLAG      = "UNFLAG"
	DIG         = "DIG"
	JOIN        = "JOIN"
)

type Action struct {
	Action string `json:"action"`
}

func (c *Client) parse(message []byte) {
	var action Action
	json.Unmarshal(message, &action)

	switch action.Action {
	case CREATE_GAME:
		c.createGame()
	}
}

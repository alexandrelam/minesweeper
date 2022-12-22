package main

const (
	CREATE_GAME = "CREATE_GAME"
	FLAG        = "FLAG"
	UNFLAG      = "UNFLAG"
	DIG         = "DIG"
	MOUSE       = "MOUSE"
)

type Action struct {
	Action string `json:"action"`
}

type ActionSquare struct {
	Action string `json:"action"`
	Row    int    `json:"row"`
	Column int    `json:"column"`
}

func (c *Client) parse(message []byte) {
	var action Action
	json.Unmarshal(message, &action)

	switch action.Action {
	case CREATE_GAME:
		c.createGame()
	case FLAG:
		var actionSquare ActionSquare
		json.Unmarshal(message, &actionSquare)

		c.flag(actionSquare.Row, actionSquare.Column)
	case UNFLAG:
		var actionSquare ActionSquare
		json.Unmarshal(message, &actionSquare)

		c.unflag(actionSquare.Row, actionSquare.Column)
	case DIG:
		var actionSquare ActionSquare
		json.Unmarshal(message, &actionSquare)

		c.dig(actionSquare.Row, actionSquare.Column)
	case MOUSE:
		c.hub.broadcast <- message
	}

}

package websocket

import "encoding/json"

const (
	CREATE_GAME = "CREATE_GAME"
	FLAG        = "FLAG"
	UNFLAG      = "UNFLAG"
	DIG         = "DIG"
	USER_MOUSE  = "USER_MOUSE"
)

type Action struct {
	Action string `json:"action"`
}

type ActionSquare struct {
	Action string `json:"action"`
	Row    int    `json:"row"`
	Column int    `json:"column"`
}

type ActionMouse struct {
	Action string `json:"action"`
	MouseX int    `json:"mouseX"`
	MouseY int    `json:"mouseY"`
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
	case USER_MOUSE:
		var actionMouse ActionMouse
		json.Unmarshal(message, &actionMouse)

		c.userMouse(actionMouse.MouseX, actionMouse.MouseY)
	}

}

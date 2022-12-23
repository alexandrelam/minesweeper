package main

const (
	GAME_LOST       = "GAME_LOST"
	UPDATE_BOARD    = "UPDATE_BOARD"
	CONNECTED_USERS = "CONNECTED_USERS"
)

type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func newReponse(responseType string, data interface{}) Response {
	return Response{
		Type: responseType,
		Data: data,
	}
}

func (r *Response) toJSON() []byte {
	jsonResponse, _ := json.Marshal(r)
	return jsonResponse
}

package main

const (
	UPDATE_BOARD    = "UPDATE_BOARD"
	CONNECTED_USERS = "CONNECTED_USERS"
	USER_MOUSE      = "USER_MOUSE"
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
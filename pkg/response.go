package main

import (
	"time"
)

const (
	GAME_LOST       = "GAME_LOST"
	GAME_WON        = "GAME_WON"
	UPDATE_BOARD    = "UPDATE_BOARD"
	HISTORY         = "HISTORY"
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

type Event struct {
	Action     string `json:"action"`
	AuthorID   string `json:"authorID"`
	AuthorName string `json:"authorName"`
	Date       string `json:"date"`
	Row        int    `json:"row"`
	Column     int    `json:"column"`
}

var NilEvent = Event{}

var history []Event = []Event{}

func newEvent(action, authorID, authorName string, row, colomn int) Event {
	return Event{
		Action:     action,
		AuthorID:   authorID,
		AuthorName: authorName,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		Row:        row,
		Column:     colomn,
	}
}

func appendEvent(event Event) {
	history = append(history, event)
}

func getHistory() []Event {
	return history
}

func resetHistory() {
	history = []Event{}
}

func (c *Client) sendHistory() {
	response := newReponse(HISTORY, history)
	c.hub.broadcast <- response.toJSON()
}

package main

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	MouseX int    `json:"mouseX"`
	MouseY int    `json:"mouseY"`
}

func GetUser(client *Client) *User {
	return &User{
		ID:     client.id,
		Name:   client.name,
		MouseX: client.mouseX,
		MouseY: client.mouseY,
	}
}

func GetAllUsers(hub *Hub) []*User {
	users := make([]*User, 0)
	for client := range hub.clients {
		users = append(users, GetUser(client))
	}

	return users
}

func (c *Client) userMouse(mouseX, mouseY int) {
	c.mouseX = mouseX
	c.mouseY = mouseY

	response := newReponse(USER_MOUSE, GetUser(c))

	c.hub.broadcast <- response.toJSON()
}

func sendUpdatedUsers(h *Hub, responseType string) {
	allUsers := GetAllUsers(h)
	response := newReponse(responseType, allUsers)

	for c := range h.clients {
		select {
		case c.send <- response.toJSON():
		default:
			close(c.send)
			delete(h.clients, c)
		}
	}
}

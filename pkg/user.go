package main

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	MouseX int    `json:"mouseX"`
	MouseY int    `json:"mouseY"`
}

func GetUser(client *Client) *User {
	return &User{
		ID:   client.id,
		Name: client.name,
	}
}

func GetAllUsers(hub *Hub) []*User {
	users := make([]*User, 0)
	for client := range hub.clients {
		users = append(users, GetUser(client))
	}

	return users
}

package ws

import "golang.org/x/net/websocket"

type Client struct {
	Conn *websocket.Conn
	RoomID string
	Message chan *Message
}
type Message struct{
	Content string
	RoomID string
	Username string
}

func (c *Client)WriteMessage() error {
	return nil
}
func (c *Client)ReadMessage() error {
	return nil
}
package service

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	id     int64
	socket *websocket.Conn
}

func (c *client) Read() {
	for {
		msgType, p, err := c.socket.ReadMessage()
		if err != nil {
			log.Printf("read user %d's message error: %v\n", c.id, err)
			break
		}
		if err := parser.Parse(msgType, p); err != nil {
			log.Printf("parse user %d's data error: %v\n", c.id, err)
			break
		}
	}

	r.Unregister <- c
}

func RegisterClient(id int64, socket *websocket.Conn) {
	r.Register <- &client{id, socket}
}

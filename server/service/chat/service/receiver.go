package service

import (
	"log"

	"github.com/Wishrem/wuso/pkg/pool"
	prot "github.com/Wishrem/wuso/pkg/protocol"
	"github.com/gorilla/websocket"
)

type Receiver struct {
	Register   chan *client
	Unregister chan *client
	clients    map[int64]*client
	pool       pool.Pool
}

// TODO
func (r Receiver) ReceiveACK(id int64) {
}

// TODO
func (r Receiver) ReceiveError(id int64, errCode prot.ErrCode) {

}

func (r Receiver) ReceiveText(id, from, to int64, text string, needACK bool) {
	go r.pool.Add(func() {
		if v, ok := r.clients[to]; ok {
			b, err := prot.BuildMsgBytes(id, from, to, text)
			if err != nil {
				log.Println("build message bytes error:", err)
				if !needACK {
					return
				}
				// TODO reply
			}
			if err := v.socket.WriteMessage(websocket.BinaryMessage, b); err != nil {
				log.Printf("write message to user %d error:%v\n", v.id, err)
				r.Unregister <- v
			}
		}
		// TODO Offline
		log.Printf("target %d is offline\n", to)
	})
}

func (r *Receiver) start() {
	for {
		select {
		case c := <-r.Register:
			go r.pool.Add(func() {
				if _, ok := r.clients[c.id]; ok {
					c.socket.Close()
				} else {
					r.clients[c.id] = c
					go c.Read()
					b, _ := prot.BuildMsgBytes(-1, -1, c.id, "connected successfully")
					if err := c.socket.WriteMessage(websocket.BinaryMessage, b); err != nil {
						log.Printf("write message to user %d error:%v\n", c.id, err)
						r.Unregister <- c
					}
				}
			})
		case c := <-r.Unregister:
			go r.pool.Add(func() {
				c.socket.Close()
				delete(r.clients, c.id)
			})
		}
	}
}

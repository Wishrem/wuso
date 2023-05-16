package message

import (
	"github.com/Wishrem/wuso/client/types"
	"github.com/Wishrem/wuso/client/websocket"
)

var m *manage

type manage struct {
	// TODO
}

func init() {
	m = &manage{}
}

func (m *manage) SendMsg(msg ...*types.Message) error {
	return nil
}

func (m *manage) RecvMsg(bytes []byte) {
	msg := new(types.Message)
	if err := msg.Unmarshal(bytes); err != nil {
		websocket.SendErr(websocket.ErrRecvWrongBytes)
	}
}

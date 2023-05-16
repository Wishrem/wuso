package message

import (
	"github.com/Wishrem/wuso/client/storage"
	"github.com/Wishrem/wuso/pkg/pool"
	"google.golang.org/genproto/googleapis/storage/v1"
)

func SendMsg(s string) error {
	// TODO
	return m.SendMsg(nil...)
}

func RecvMsg(p pool.Param) {
	var msg Message
	if m, ok := p["msg"]; !ok {
		return
	} else if msg, ok = m.(Message); !ok {
		return
	}

	// Storage
	if err := storage.SaveMsg(msg); err != nil {

	}

	// ACK Server

	// Display
}

func GetMsgs() ([]string, error) {
	return nil, nil
}

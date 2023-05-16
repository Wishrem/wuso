package link

import (
	"github.com/Wishrem/wuso/client/message"
	"github.com/Wishrem/wuso/pkg/pool"
)

var MsgPool *pool.Pool

func RecvMsg(msg []byte) {
	m := new(message.Message)
	if err := m.Unmarshal(msg); err != nil {
		return
	}

	MsgPool.AddJob(pool.NewJob(pool.Param{
		"msg": m,
	}, message.RecvMsg))
}

package service

import (
	"github.com/Wishrem/wuso/config"
	"github.com/Wishrem/wuso/pkg/pool"
	prot "github.com/Wishrem/wuso/pkg/protocol"
)

var parser *prot.Parser
var r *Receiver

func Init(worker int) {
	r = &Receiver{
		Register:   make(chan *client, worker/2),
		Unregister: make(chan *client, worker/2),
		clients:    make(map[int64]*client),
		pool:       *pool.NewPool(worker),
	}

	parser = prot.NewParser(r, config.Server.WithoutClient)

	go r.start()
}

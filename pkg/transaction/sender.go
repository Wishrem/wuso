package transaction

import (
	"sync"

	"github.com/Wishrem/wuso/pkg/pool"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/eapache/queue.v1"
)

type sender struct {
	txId chan int64
	pool *pool.Pool
	q    *waitQue
}

type waitQue struct {
	*sync.RWMutex
	*queue.Queue
}

func newWaitQue() *waitQue {
	return &waitQue{
		new(sync.RWMutex),
		queue.New(),
	}
}

func newSender(workers int) *sender {
	return &sender{
		txId: make(chan int64, workers>>1),
		pool: pool.NewPool(workers),
		q:    newWaitQue(),
	}
}

func (s *sender) makeParam() *pool.Param {
	return &pool.Param{
		"q": *s.q,
	}
}

var json = jsoniter.ConfigFastest

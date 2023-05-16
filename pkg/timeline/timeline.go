package timeline

import (
	"container/heap"
	"sync"
	"time"

	"gopkg.in/eapache/queue.v1"
)

type element struct {
	data        interface{}
	unixMillion int64
}

type recvQue struct {
	*sync.Mutex
	*queue.Queue
}

type readCond struct {
	*sync.Cond
	mutex *sync.Mutex
}

type TimeLine struct {
	rq   *recvQue
	heap *bufHeap
	read *readCond
}

func newRecvQue() *recvQue {
	return &recvQue{
		new(sync.Mutex),
		queue.New(),
	}
}

func newReadCond() *readCond {
	mutex := new(sync.Mutex)
	return &readCond{
		sync.NewCond(mutex),
		mutex,
	}
}

func New() *TimeLine {
	return &TimeLine{
		newRecvQue(),
		newBufHeap(),
		newReadCond(),
	}
}

func (tl *TimeLine) Add(data interface{}, unixMillion int64) {
	tl.heap.Lock()
	heap.Push(tl.heap, &element{
		data:        data,
		unixMillion: unixMillion,
	})
	tl.heap.Unlock()

	time.Sleep(time.Second)
	tl.heap.Lock()
	defer tl.heap.Unlock()
	tl.rq.Lock()
	defer tl.rq.Unlock()
	for tl.heap.Len() != 0 {
		tl.rq.Add(heap.Pop(tl.heap))
	}
	tl.read.Broadcast()
}

func (tl *TimeLine) Get(res chan interface{}) {
	for {
		tl.read.L.Lock()
		tl.read.Wait()

		tl.rq.Lock()
		for tl.rq.Length() != 0 {
			res <- tl.rq.Remove()
		}
		tl.rq.Unlock()
		tl.read.L.Unlock()
	}
}

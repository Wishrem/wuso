package pool

import (
	"sync"

	"gopkg.in/eapache/queue.v1"
)

type worker struct {
	id   int
	job  chan func()
	done chan int
}

func (w *worker) start() {
	for {
		job := <-w.job
		job()
		w.done <- w.id
	}
}

type accessibleQueue struct {
	*sync.Mutex
	*queue.Queue
}

type Pool struct {
	workers []*worker
	jobs    chan func()
	aq      accessibleQueue
}

func NewPool(workers int) *Pool {
	p := &Pool{
		make([]*worker, workers),
		make(chan func(), workers),
		accessibleQueue{
			new(sync.Mutex),
			queue.New(),
		},
	}
	done := make(chan int, workers)
	for i := 0; i < workers; i++ {
		worker := &worker{i, make(chan func(), 1), done}
		p.workers[i] = worker
		p.aq.Add(i)
		go worker.start()
	}

	go func() {
		for {
			select {
			case id := <-done:
				p.aq.Lock()
				p.aq.Add(id)
				p.aq.Unlock()

			case job := <-p.jobs:
				p.aq.Lock()
				id := p.aq.Remove().(int)
				p.aq.Unlock()
				p.workers[id].job <- job
			}
		}
	}()

	return p
}

func (p *Pool) Add(job func()) {
	p.jobs <- job
}

package timeline

import "sync"

type BufHeap *bufHeap

type bufHeap struct {
	*sync.Mutex
	data []*element
}

func newBufHeap() *bufHeap {
	return &bufHeap{
		new(sync.Mutex),
		make([]*element, 0),
	}
}

func (bh *bufHeap) Len() int {
	return len(bh.data)
}

func (bh *bufHeap) Less(i, j int) bool {
	return bh.data[i].unixMillion < bh.data[j].unixMillion
}

func (bh *bufHeap) Swap(i, j int) {
	bh.data[i], bh.data[j] = bh.data[j], bh.data[i]
}

func (bh *bufHeap) Push(x interface{}) {
	bh.data = append(bh.data, x.(*element))
}

func (bh *bufHeap) Pop() interface{} {
	x := bh.data[len(bh.data)-1]
	bh.data = bh.data[:len(bh.data)-1]
	return x
}

package message

import "sync"

type Buffer struct {
	requestQueue  []*Request
	requestLocker *sync.RWMutex
}

func NewBuffer() *Buffer {
	return &Buffer{
		requestQueue:  make([]*Request, 0),
		requestLocker: new(sync.RWMutex),
	}
}

func (b *Buffer) AppendToRequestQueue(req *Request) {
	b.requestLocker.Lock()
	b.requestQueue = append(b.requestQueue, req)
	b.requestLocker.Unlock()
}

func (b *Buffer) RequestQueueEmpty() bool {
	ret := false
	b.requestLocker.RLock()

	b.requestLocker.RUnlock()
	return ret
}

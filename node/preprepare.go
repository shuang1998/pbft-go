package node

import "time"

// send pre-prepare thread by request notify or timer
func (n *Node)PrePrepareSendThread() {
	// TODO change timer duration form config
	duration := time.Second
	timer := time.After(duration)
	for {
		select {
		// recv request or time out
		case <-n.prePrepareNotify:
			n.PrePrepareSendHandleFunc()
		case <-timer:
			timer = nil
			timer = time.After(duration)
		}
	}
}

func (n *Node)PrePrepareSendHandleFunc() {
	// buffer is empty or execute op num max
	if n.buffer.RequestQueueEmpty() || n.executeNum.Get() >= n.cfg.ExecuteMaxNum {
		return
	}

}


// recv pre-prepare thread from http server
func (n *Node) PrePrepareRecvThread() {

}

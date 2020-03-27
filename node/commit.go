package node

import (
	"github.com/pipapa/pbft/message"
)

func (n *Node) commitRecvThread() {
	for {
		select {
		case msg := <-n.commit:
			if !n.checkCommitMsg(msg) {
				continue
			}
			// buffer the commit msg
			n.buffer.BufferCommitMsg(msg)
			if n.buffer.IsReadyToExecute(msg.Digest, n.cfg.FaultNum, msg.View, msg.Sequence) {
				// buffer to ExcuteQueue
				n.buffer.AppendToExecuteQueue(msg)
				// notify ExcuteThread
				n.executeNotify<-true
			}
		}
	}
}

func (n *Node) checkCommitMsg(msg *message.Commit) bool {
	if n.view != msg.View {
		return false
	}
	if !n.sequence.CheckBound(msg.Sequence) {
		return false
	}
	return true
}
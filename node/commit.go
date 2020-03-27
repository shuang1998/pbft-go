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
			if n.buffer.IsTrueOfCommitMsg(msg.Digest, n.cfg.FaultNum) &&
			   n.buffer.IsTrueOfPrepareMsg(msg.Digest, n.cfg.FaultNum) &&
			   n.buffer.IsExistPreprepareMsg(msg.View, msg.Sequence) {
				// clear the buffer about msg.Digest
				n.buffer.ClearCommitMsg(msg.Digest)
				n.buffer.ClearPrepareMsg(msg.Digest)
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
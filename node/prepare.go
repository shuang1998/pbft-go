package node

import (
	"github.com/pipapa/pbft/message"
	"github.com/pipapa/pbft/server"
)

func (n *Node) prepareRecvAndCommitSendThread() {
	for {
		select {
		case msg := <-n.prepareRecv:
			if !n.checkPrepareMsg(msg) {
				continue
			}
			// buffer the prepare msg
			n.buffer.BufferPrepareMsg(msg)
			// verify send commit msg
			if n.buffer.IsTrueOfPrepareMsg(msg.Digest, n.cfg.FaultNum) {
				content, msg, err := message.NewCommitMsg(n.id, msg)
				if err != nil {
					continue
				}
				// buffer commit msg
				n.buffer.BufferCommitMsg(msg)
				// TODO broadcast error when buffer the commit msg
				n.BroadCast(content, server.CommitEntry)
			}
		}
	}
}

func (n *Node) checkPrepareMsg(msg *message.Prepare) bool {
	if n.view != msg.View {
		return false
	}
	if !n.sequence.CheckBound(msg.Sequence) {
		return false
	}
	return true
}
package node

import (
	"github.com/pipapa/pbft/message"
	"log"
)

func (n *Node) executeAndReplyThread() {
	for {
		select {
		case <-n.executeNotify:
			// execute batch
			batchs, lastSeq := n.buffer.BatchExecute(n.sequence.GetLastSequence())
			n.sequence.SetLastSequence(lastSeq)
			// map the digest to request
			requestBatchs := make([]*message.Request, 0)
			for _, b := range batchs {
				reqs := n.buffer.FetchRequest(b)
				requestBatchs = append(requestBatchs, reqs...)
			}
			// execute the request
			for _, r := range requestBatchs {
				log.Printf("do the opreation %s - %d", r.Op, r.TimeStamp)
			}

		}
	}
}

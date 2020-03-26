package node

import (
	"github.com/pipapa/pbft/cmd"
	"github.com/pipapa/pbft/message"
	"github.com/pipapa/pbft/server"
	"log"
)

type Node struct {
	cfg    *cmd.SharedConfig
	server *server.HttpServer

	id       message.Identify
	view     message.View
	table    map[message.Identify]string
	faultNum uint

	lastReply      *message.LastReply
	sequence       *Sequence
	executeNum     *ExecuteOpNum

	buffer         *message.Buffer

	requestRecv    chan *message.Request
	prePrepareRecv chan *message.PrePrepare
	prepareRecv    chan *message.Prepare
	commit         chan *message.Commit

	prePrepareSendNotify chan bool
	executeNotify        chan bool
}

func NewNode(cfg *cmd.SharedConfig) *Node {
	node := &Node{
		// config
		cfg:	  cfg,
		// http server
		server:   server.NewServer(cfg),
		// information about node
		id:       cfg.Id,
		view:     cfg.View,
		table:	  cfg.Table,
		faultNum: cfg.FaultNum,
		// lastReply state
		lastReply:  message.NewLastReply(),
		sequence:   NewSequence(cfg),
		executeNum: NewExecuteOpNum(),
		// the message buffer to store msg
		buffer: message.NewBuffer(),
		// chan for server and recv thread
		requestRecv:    make(chan *message.Request),
		prePrepareRecv: make(chan *message.PrePrepare),
		prepareRecv:    make(chan *message.Prepare),
		commit:         make(chan *message.Commit),
		// chan for notify pre-prepare send thread
		prePrepareSendNotify: make(chan bool),
		// chan for notify execute op and reply thread
		executeNotify:        make(chan bool),
	}
	log.Printf("[Node] the node id:%d, view:%d, fault number:%d\n", node.id, node.view, node.faultNum)
	return node
}

func (n *Node) Run() {
	// first register chan for server
	n.server.RegisterChan(n.requestRecv, n.prePrepareRecv, n.prepareRecv, n.commit)
	go n.server.Run()
	go n.requestRecvThread()
	go n.prePrepareSendThread()
	go n.prePrepareRecvAndPrepareSendThread()
	go n.prepareRecvAndCommitSendThread()
	go n.commitRecvThread()
	go n.executeAndReplyThread()
}

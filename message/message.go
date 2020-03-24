package message

type TimeStamp uint64	// 时间戳格式
type Identify  uint64   // 客户端标识格式
type View      uint64	// 视图
type Sequence  uint64	// 序号

// Operation
type Operation struct {

}
// Result
type Result struct {

}
// Request
type Request struct {
	Op        Operation `json:"operation"`
	TimeStamp TimeStamp `json:"timestamp"`
	ID        Identify	`json:"clientID"`
	Digest    string	`json:"digest"`
}
// Message
type Message struct {
	Requests []*Request `json:"requests"`
}
// Pre-Prepare
type PrePrepare struct {
	View      View		`json:"view"`
	Sequence  Sequence	`json:"sequence"`
	Digest    string	`json:"digest"`
	Message   Message	`json:"message"`
}
// Prepare
type Prepare struct {
	View	  View	    `json:"view"`
	Sequence  Sequence  `json:"sequence"`
	Digest    string	`json:"digest"`
	Identify  Identify	`json:"id"`
}
// Commit
type Commit struct {
	View	  View		`json:"view"`
	Sequence  Sequence	`json:"sequence"`
	Digest    string	`json:"digest"`
	Identify  Identify	`json:"id"`
}
// Reply
type Reply struct {
	View	  View		`json:"view"`
	TimeStamp TimeStamp `json:"timestamp"`
	Id        Identify  `json:"nodeID"`
	Result    Result	`json:"result"`
}

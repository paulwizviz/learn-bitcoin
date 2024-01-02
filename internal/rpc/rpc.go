package rpc

const (
	GetNewAddress = "getnewaddress"
)

type Body struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      string `json:"id"`
}

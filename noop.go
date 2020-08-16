package mdp

// NoOp represents a yet to be defined runtime operation
type NoOp struct {
	Identifier
	Arguments interface{} `json:"command"`
}

func (no *NoOp) String() string {
	return "noop"
}

// Code for NoOp
func (no *NoOp) Code() byte {
	return 0
}

// Bytes for NoOp
func (no *NoOp) Bytes() (bs []byte) {
	bs = make([]byte, 0, 1)
	bs = append(bs, 0)
	return FormatBytes(bs)
}

// Validate nothing
func (no *NoOp) Validate() (err error) {
	return
}

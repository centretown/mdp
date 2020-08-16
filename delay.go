package mdp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DelayOp data required for delay operaton
type DelayOp struct {
	Identifier
	Command struct {
		Duration uint16 `json:"duration"`
	} `json:"command"`
}

// Code unique identifying byte
func (dly *DelayOp) Code() byte {
	return 'D'
}

func (dly *DelayOp) String() string {
	return fmt.Sprintf("delay %d ms", dly.Command.Duration)
}

// Bytes -
func (dly *DelayOp) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, Endian, dly.Code())
	binary.Write(buf, Endian, dly.Command.Duration)
	return FormatBytes(buf.Bytes())
}

// Validate -
func (dly *DelayOp) Validate() (err error) {
	return
}

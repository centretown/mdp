package mdp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// HallOp -
type HallOp struct {
	Identifier
	Command struct {
		Measurement uint32 `json:"measurement"`
	} `json:"command"`
}

func (h *HallOp) String() string {
	return fmt.Sprintf("read hall measurement")
}

// Code one byte id
func (h *HallOp) Code() byte {
	return 'H'
}

// Bytes -
func (h *HallOp) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, Endian, h.Code())
	binary.Write(buf, Endian, h.Command.Measurement)
	return FormatBytes(buf.Bytes())
}

// Validate -
func (h *HallOp) Validate() (err error) {
	return
}

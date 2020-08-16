package mdp

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// Endian -
var Endian = binary.LittleEndian

// Identifier -
type Identifier struct {
	Sequence int32  `json:"sequence"`
	Token    string `json:"type"`
}

// Validator checks itself and returns an error if invalid
type Validator interface {
	Validate() error
}

// CodeWriter generates code and descriptive strings
type CodeWriter interface {
	Code() byte
	Bytes() []byte
	fmt.Stringer
}

// Operator validates and writes code
type Operator interface {
	Validator
	CodeWriter
}

// Operation -
type Operation struct {
	Operator
}

// UnmarshalJSON to defined operation and validate
func (op *Operation) UnmarshalJSON(b []byte) (err error) {
	// use noop for inspection
	noop := NoOp{}
	err = json.Unmarshal(b, &noop)
	if err != nil {
		return
	}

	// commit
	switch noop.Token {
	case "DELAY":
		op.Operator = &DelayOp{}
	case "MODE":
		op.Operator = &ModeOp{}
	case "PIN":
		op.Operator = &PinOp{}
	case "HALL":
		op.Operator = &HallOp{}
	default:
		err = fmt.Errorf("unknown action type")
		return
	}

	err = json.Unmarshal(b, &op.Operator)
	if err != nil {
		fmt.Println(err)
	}

	err = op.Operator.Validate()
	return
}

// FormatBytes returns input bytes preceded by the byte length
func FormatBytes(buf []byte) (bs []byte) {
	bl := byte(len(buf) + 1)
	bs = make([]byte, 0, bl)
	bs = append(bs, bl)
	bs = append(bs, buf...)
	return
}

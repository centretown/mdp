package mdp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// ModeArguments provide input to the pinMode command
type ModeArguments struct {
	Pin    int8   `json:"id"`
	Signal string `json:"signal"`
	Mode   string `json:"mode"`
}

// ModeOp data required for mode operaton
type ModeOp struct {
	Identifier
	Arguments ModeArguments `json:"command"`
}

func (m *ModeOp) String() string {
	return fmt.Sprintf("prepare pin #%d for %s %s",
		m.Arguments.Pin, m.Arguments.Signal, m.Arguments.Mode)
}

// Code one byte id
func (m *ModeOp) Code() byte {
	return 'M'
}

// Bytes -
func (m *ModeOp) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, Endian, m.Code())
	binary.Write(buf, Endian, m.Arguments.Pin)
	binary.Write(buf, Endian, []byte(m.Arguments.Signal[:1]))
	binary.Write(buf, Endian, []byte(m.Arguments.Mode[:1]))
	return FormatBytes(buf.Bytes())
}

var (
	// ErrSignal unexpected signal
	ErrSignal = errors.New("want 'analog' or 'digital'")
	// ErrMode unexpected mode
	ErrMode = errors.New("want 'input' or 'output'")
	// ErrPin unexpected pin id
	ErrPin = errors.New("want pin # '0-63'")
	// ErrParseMode unable to parse Mode command
	ErrParseMode = errors.New("Mode parsing errors")
)

// Validate mode arguments
func (m *ModeOp) Validate() (err error) {
	err = ValidateModeArguments(m.Sequence, m, &m.Arguments)
	if err != nil {
		err = WrapInfo(err, ErrParseMode, m.Sequence, m)
	}
	return
}

// ValidateModeArguments reports any errors
func ValidateModeArguments(sequence int32, op Operator, args *ModeArguments) (err error) {
	switch args.Signal {
	case "analog":
	case "digital":
	default:
		err = WrapInfo(err, ErrSignal, sequence, args.Signal)
	}

	switch args.Mode {
	case "input":
	case "output":
	default:
		err = WrapInfo(err, ErrMode, sequence, args.Mode)
	}

	if args.Pin < 0 || args.Pin > 63 {
		err = WrapInfo(err, ErrPin, sequence, args.Pin)
	}
	return

}

package mdp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// PinArguments defines arguments to Pin Operations
type PinArguments struct {
	ModeArguments
	Value int16 `json:"value"`
}

// PinOp data required for mode operaton
type PinOp struct {
	Identifier
	Arguments PinArguments `json:"command"`
}

// Code one byte id
func (p *PinOp) Code() byte {
	return 'P'
}

func (p *PinOp) String() string {
	if p.Arguments.Mode == "output" {
		return fmt.Sprintf("set pin #%d to %s %d",
			p.Arguments.Pin, p.Arguments.Signal, p.Arguments.Value)
	}
	return fmt.Sprintf("read %s from pin #%d",
		p.Arguments.Signal, p.Arguments.Pin)
}

// Bytes encapsulates length 1 byte, code 1 byte, byte arguments...
func (p *PinOp) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, Endian, p.Code())
	binary.Write(buf, Endian, p.Arguments.Pin)
	binary.Write(buf, Endian, []byte(p.Arguments.Signal[:1]))
	binary.Write(buf, Endian, []byte(p.Arguments.Mode[:1]))
	binary.Write(buf, Endian, p.Arguments.Value)
	return FormatBytes(buf.Bytes())
}

var (
	// ErrPinValue unexpected pin value
	ErrPinValue = errors.New("want pin value '0' or '1'")
	// ErrParsePin unable to parse Pin command
	ErrParsePin = errors.New("Pin parsing errors")
)

// Validate -
func (p *PinOp) Validate() (err error) {
	err = ValidateModeArguments(p.Sequence, p, &p.Arguments.ModeArguments)

	if p.Arguments.Mode == "output" &&
		p.Arguments.Signal == "digital" &&
		(p.Arguments.Value < 0 || p.Arguments.Value > 1) {

		err = WrapInfo(err, ErrPinValue, p.Sequence, p.Arguments.Value)
	}

	if err != nil {
		err = WrapInfo(err, ErrParsePin, p.Sequence, p)
	}

	return
}

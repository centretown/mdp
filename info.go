package mdp

import (
	"fmt"
)

// Info wraps an Operator and Sequence with the error
type Info struct {
	Sequence int32
	Text     interface{}
	Kind     error
	Err      error
}

var sep = "\n"

func (e *Info) Error() string {
	return fmt.Sprintf("\n%d: have '%v' %v", e.Sequence, e.Text, e.Err)
}

// Unwrap a parsing error
func (e *Info) Unwrap() error {
	return e.Err
}

// Is this error that kind
func (e *Info) Is(kind error) bool {
	return e.Kind == kind
}

// WrapInfo wraps existing and id the new ones' kind
func WrapInfo(current, wrap error, seq int32, text interface{}) error {
	wrapped := Info{Sequence: seq, Text: text, Kind: wrap}
	if current != nil {
		wrap = fmt.Errorf("%v %w", wrap, current)
	}
	wrapped.Err = wrap
	return &wrapped
}

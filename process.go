package mdp

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Process - describes an arduino based sketch
type Process struct {
	Label     string      `json:"label"`
	DeviceKey string      `json:"deviceKey"`
	Purpose   string      `json:"purpose"`
	Setup     []Operation `json:"setup"`
	Loop      []Operation `json:"loop"`
}

// ErrMakeProcess -
var ErrMakeProcess = errors.New("Make Process")

// MakeProcess data from byte stream
func MakeProcess(data []byte) (process *Process, err error) {
	process = &Process{}
	err = json.Unmarshal(data, process)
	if err != nil {
		err = WrapInfo(err, ErrMakeProcess, 0, "json.Unmarshall")
	}
	return
}

// ErrMakeProcessFromFile -
var ErrMakeProcessFromFile = errors.New("Make Process From File")

// MakeProcessFromFile name
func MakeProcessFromFile(fname string) (process *Process, err error) {
	var data []byte
	data, err = ioutil.ReadFile(fname)
	if err != nil {
		err = WrapInfo(err, ErrMakeProcessFromFile, 0, fname)
		return
	}
	process, err = MakeProcess(data)
	if err != nil {
		err = WrapInfo(err, ErrMakeProcessFromFile, 0, fname)
	}
	return
}

package mdp

import (
	"errors"
	"testing"
)

func TestValidProcess(t *testing.T) {
	process, err := MakeProcessFromFile("test_data/process.json")
	if err != nil {
		t.Fatalf("while making a process: %v", err)
		return
	}
	showOps(t, process.Setup)
	showOps(t, process.Loop)
}

func TestBadMode(t *testing.T) {
	_, err := MakeProcessFromFile("test_data/bad_mode.json")
	if err == nil {
		t.Fatalf("error undetected: %v", err)
	}

	if !errors.Is(err, ErrSignal) {
		t.Fatalf("signal error undetected")
	} else {
		t.Logf("signal error detected")
	}
	if !errors.Is(err, ErrMode) {
		t.Fatalf("mode error undetected")
	} else {
		t.Logf("mode error detected")
	}
	if !errors.Is(err, ErrPin) {
		t.Fatalf("pin error undetected")
	} else {
		t.Logf("pin error detected")
	}
	t.Logf("%v", err)

}
func TestBadPin(t *testing.T) {
	_, err := MakeProcessFromFile("test_data/bad_pin.json")
	if err == nil {
		t.Fatalf("error undetected: %v", err)
	}

	if errors.Is(err, ErrSignal) {
		t.Fatalf("signal error detected")
	}
	if errors.Is(err, ErrMode) {
		t.Fatalf("mode error detected")
	}
	if errors.Is(err, ErrPin) {
		t.Log("pin error detected")
	} else {
		t.Fatalf("pin error undetected")
	}
	if errors.Is(err, ErrPinValue) {
		t.Log("pin value error detected")
	} else {
		t.Fatalf("pin value error undetected")
	}

	t.Logf("%v", err)

}

func showOps(t *testing.T, p []Operation) {
	t.Log(p)
	for _, r := range p {
		t.Log(r)
		t.Logf("%x", r.Operator.Bytes())
	}
}

func TestBadJson(t *testing.T) {
	_, err := MakeProcessFromFile("test_data/bad.json")
	if err == nil {
		t.Fatal("Bad Json expected")
	}
	t.Logf("%v", err)
}

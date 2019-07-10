package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"testing"
)

func TestUnknownStruct(t *testing.T) {
	state := states.ScooterUnknown{}
	ret, err := state.Next()

	if err == nil {
		t.Fatalf("ScooterUnknown expected to not have next function %v", ret)
	}

}

func TestScooterUnknownIsValid(t *testing.T) {
	state := states.ScooterUnknown{}
	ret, _ := state.IsValid()

	if ret != false {
		t.Fatalf("ScooterUnknown must not have implementation")
	}

}

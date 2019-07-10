package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"testing"
)

func TestStateStruct(t *testing.T) {
	state := states.ScooterState{}
	ret, err := state.Next()

	if err == nil {
		t.Fatalf("State expected to not have next function %v", ret)
	}

}

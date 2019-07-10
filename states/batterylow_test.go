package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"reflect"
	"testing"
	"time"
)

func TestGoBounty(t *testing.T) {
	state := states.ScooterBatteryLow{}
	state.Name = "test-name"
	state.User = nil
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterBounty{})) {
		t.Fatalf("Expected Bounty, found %v", reflect.TypeOf(newstate))
	}
}

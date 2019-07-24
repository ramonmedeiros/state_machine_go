package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"testing"
	"time"
)

func TestGoBounty(t *testing.T) {
	state := states.ScooterBatteryLow{}
	state.User = nil
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterBounty{})) {
		t.Fatalf("Expected Bounty, found %v", reflect.TypeOf(newstate))
	}
}

func TestValidScooterBatteryLow(t *testing.T) {
	state := states.ScooterBatteryLow{}
	state.User = nil
	state.BatteryLevel = 21
	state.LastStateChange = time.Now()

	newstate, _ := state.IsValid()

	if newstate != false {
		t.Fatalf("Expected invalid state, not valid")
	}
}

func TestScooterBatteryLowWithUser(t *testing.T) {
	user := users.User{}
	state := states.ScooterBatteryLow{}
	state.User = user
	state.BatteryLevel = 19

	newstate, _ := state.IsValid()

	if newstate != false {
		t.Fatalf("Expected invalid state, not valid")
	}
}

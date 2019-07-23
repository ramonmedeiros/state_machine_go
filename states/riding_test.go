package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"testing"
	"time"
)

func TestKeepRiding(t *testing.T) {
	user := users.User{}
	state := states.ScooterRiding{}
	state.User = &user
	state.BatteryLevel = 100
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterRiding{})) {
		t.Fatalf("Expected Riding, found %v", reflect.TypeOf(newstate))
	}
}

func TestUserDeattached(t *testing.T) {
	state := states.ScooterRiding{}
	state.User = nil
	state.BatteryLevel = 100
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterReady{})) {
		t.Fatalf("Expected Ready, found %v", reflect.TypeOf(newstate))
	}
}

func TestBatteryLow(t *testing.T) {
	user := users.User{}
	state := states.ScooterRiding{}
	state.User = &user
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterBatteryLow{})) {
		t.Fatalf("Expected BatteryLow, found %v", reflect.TypeOf(newstate))
	}
}

func TestBatteryLowNoUser(t *testing.T) {
	state := states.ScooterRiding{}
	state.User = nil
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterBatteryLow{})) {
		t.Fatalf("Expected BatteryLow, found %v", reflect.TypeOf(newstate))
	}
}

func TestRidingInvalidLowBattery(t *testing.T) {
	user := users.User{}
	state := states.ScooterRiding{}
	state.User = user
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	status, _ := state.IsValid()

	if status != false {
		t.Fatalf("Riding cannot have low battery")
	}
}

func TestRidingInvalidUser(t *testing.T) {
	state := states.ScooterRiding{}
	state.User = nil
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	status, _ := state.IsValid()

	if status != false {
		t.Fatalf("Riding should have associated user")
	}
}

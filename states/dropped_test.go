package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"testing"
	"time"
)

func TestDroppedNoUser(t *testing.T) {
	state := states.ScooterDropped{}
	state.User = nil
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if newstate != false {
		t.Fatalf("Expected failure, got %v", newstate)
	}
}

func TestDroppedNormalUser(t *testing.T) {
	user := users.User{}
	state := states.ScooterDropped{}
	state.User = &user
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if newstate != false {
		t.Fatalf("Expected failure, got %v", newstate)
	}
}

func TestDroppedToReady(t *testing.T) {
	user := users.Hunter{}
	state := states.ScooterDropped{}
	state.User = &user
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	// convert interface to ScooterReady
	readyState, _ := newstate.(states.ScooterReady)

	if (reflect.TypeOf(readyState) != reflect.TypeOf(states.ScooterReady{})) {
		t.Fatalf("Expected Dropped, got %v", readyState)
	}

	if readyState.BatteryLevel != 100 {
		t.Fatalf("Expected BaterryFull, got %v", readyState.BatteryLevel)
	}
}

func TestScooterDroppedValidUserUser(t *testing.T) {
	user := users.User{}

	state := states.ScooterDropped{}
	state.User = user

	ret, _ := state.Next()
	if ret != false {
		t.Fatalf("users.User expected to NOT be allowed")
	}
}

func TestScooterDroppedValidUserHunter(t *testing.T) {
	hunter := users.Hunter{}

	state := states.ScooterDropped{}
	state.User = &hunter

	ret, msg := state.Next()
	if ret == false {
		t.Fatalf("%v", msg)
	}
}

func TestScooterDroppedValidUserAdmin(t *testing.T) {
	admin := users.Admin{}

	state := states.ScooterDropped{}
	state.User = admin

    ret, msg := state.Next()
    if ret != false {
        t.Fatalf("%v", msg)
    }
}

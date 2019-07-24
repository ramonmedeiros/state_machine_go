package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"testing"
	"time"
)

func TestCollectedNoUser(t *testing.T) {
	state := states.ScooterCollected{}
	state.User = nil
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if newstate != false {
		t.Fatalf("Expected failure, got %v", newstate)
	}
}

func TestCollectedNormalUser(t *testing.T) {
	user := users.User{}
	state := states.ScooterCollected{}
	state.User = &user
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if newstate != false {
		t.Fatalf("Expected failure, got %v", newstate)
	}
}

func TestCollectedToDropped(t *testing.T) {
	user := users.Hunter{}
	state := states.ScooterCollected{}
	state.User = &user
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterDropped{})) {
		t.Fatalf("Expected Dropped, got %v", newstate)
	}
}

func TestScooterCollectedValidUserUser(t *testing.T) {
	user := users.User{}

	state := states.ScooterCollected{}
	state.User = user

	ret, _ := state.Next()
	if ret != false {
		t.Fatalf("users.User expected to NOT be allowed")
	}
}

func TestScooterCollectedValidUserHunter(t *testing.T) {
	hunter := users.Hunter{}

	state := states.ScooterCollected{}
	state.User = &hunter

	ret, msg := state.Next()
	if ret == false {
		t.Fatalf("%v", msg)
	}
}

func TestScooterCollectedValidUserAdmin(t *testing.T) {
	admin := users.Admin{}

	state := states.ScooterCollected{}
	state.User = admin

	ret, _ := state.Next()
	if ret != false {
		t.Fatalf("users.Admin expected to NOT be allowed")
	}
}

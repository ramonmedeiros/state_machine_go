package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"testing"
	"time"
)

func mockTime(hour int, minute int) {
	states.Now = func() time.Time {
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, 1, 0, time.Now().Location())
	}
}

func unmockTime() {
	states.Now = func() time.Time {
		return time.Now()
	}
}

func TestReadyToUnknownState(t *testing.T) {
	twodaysago := time.Now().Add(time.Hour * -48)
	state := states.ScooterReady{}
	state.User = nil
	state.BatteryLevel = 100
	state.LastStateChange = twodaysago

	newstate, _ := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterUnknown{})) {
		t.Fatalf("Expected unknow, found %v", reflect.TypeOf(newstate))
	}
}

func TestBountyState(t *testing.T) {
	state := states.ScooterReady{}
	state.User = nil
	state.BatteryLevel = 100
	state.LastStateChange = time.Now()

	// mock time to 21:30
	mockTime(21, 30)
	newstate, _ := state.Next()
	unmockTime()

	// rollback mock
	states.Now = time.Now

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterBounty{})) {
		t.Fatalf("Expected Bounty, found %v", reflect.TypeOf(newstate))
	}
}

func TestKeepReady(t *testing.T) {
	state := states.ScooterReady{}
	state.User = nil
	state.BatteryLevel = 100
	state.LastStateChange = time.Now()

	// mock time to avoid bounty
	mockTime(21, 00)
	newstate, _ := state.Next()
	unmockTime()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(&states.ScooterReady{})) {
		t.Fatalf("Expected Ready, found %v", reflect.TypeOf(newstate))
	}
}

func TestGoRiding(t *testing.T) {
	user := users.User{}
	state := states.ScooterReady{}
	state.User = &user
	state.BatteryLevel = 100
	state.LastStateChange = time.Now()

	// mock time to avoid bounty
	mockTime(21, 00)
	newstate, _ := state.Next()
	unmockTime()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterRiding{})) {
		t.Fatalf("Expected Riding, found %v", reflect.TypeOf(newstate))
	}
}

func TestScooterReadyValidUserUser(t *testing.T) {
	user := users.User{}

	state := states.ScooterReady{}
	state.User = user

	ret, _ := state.Next()
	if ret == false {
		t.Fatalf("users.User expected to be allowed")
	}
}

func TestScooterReadyValidUserHunter(t *testing.T) {
	hunter := users.Hunter{}

	state := states.ScooterReady{}
	state.User = hunter

	ret, _ := state.Next()
	if ret == false {
		t.Fatalf("users.Hunter expected to be allowed")
	}
}

func TestScooterReadyValidUserAdmin(t *testing.T) {
	admin := users.Admin{}

	state := states.ScooterReady{}
	state.User = admin

	ret, _ := state.Next()
	if ret == false {
		t.Fatalf("users.Admin expected to be allowed")
	}
}

func TestReadyInvalidUnknown(t *testing.T) {
	twodaysago := time.Now().Add(time.Hour * -58)
	state := states.ScooterReady{}
	state.User = nil
	state.BatteryLevel = 100
	state.LastStateChange = twodaysago

	status, _ := state.IsValid()

	if status != false {
		t.Fatalf("State is invalid: must be unknown")
	}
}

func TestReadyInvalidWithUser(t *testing.T) {
	user := users.User{}
	state := states.ScooterReady{}
	state.User = user
	state.BatteryLevel = 100
	state.LastStateChange = time.Now()

	// mock time to avoid bounty
	mockTime(21, 30)
	status, _ := state.IsValid()
	unmockTime()

	if status != false {
		t.Fatalf("Ready state cannot has user")
	}
}

func TestReadyValid(t *testing.T) {
	state := states.ScooterReady{}
	state.User = nil
	state.BatteryLevel = 100
	state.LastStateChange = time.Now()

	// mock time to avoid bounty
	mockTime(21, 00)
	status, msg := state.IsValid()
	unmockTime()

	if status != true {
		t.Fatalf("%v", msg)
	}
}

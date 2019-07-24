package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"testing"
	"time"
)

func TestBountyNoUser(t *testing.T) {
	state := states.ScooterBounty{}
	state.User = nil
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.IsValid()

	if newstate != true {
		t.Fatalf("Expected valid state")
	}
}

func TestBountyNormalUser(t *testing.T) {
	user := users.User{}
	state := states.ScooterBounty{}
	state.User = user
	state.BatteryLevel = 19
	state.LastStateChange = time.Now()

	newstate, _ := state.IsValid()

	if newstate != false {
		t.Fatalf("Expected failure, got %v", newstate)
	}
}

func TestBountyToCollected(t *testing.T) {
	user := users.Hunter{}
	state := states.ScooterBounty{}
	state.User = user

	newstate, msg := state.Next()

	if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterCollected{})) {
		t.Fatalf("%v", msg)
	}
}

func TestScooterBountyValidUserUser(t *testing.T) {
	user := users.User{}

	state := states.ScooterBounty{}
	state.User = user

	ret, _ := state.AllowedUser()
	if ret != false {
		t.Fatalf("users.User expected to NOT be allowed")
	}
}

func TestScooterBountyValidUserHunter(t *testing.T) {
	hunter := users.Hunter{}

	state := states.ScooterBounty{}
	state.User = hunter

	ret, _ := state.AllowedUser()
	if ret == false {
		t.Fatalf("users.Hunter expected to be allowed")
	}
}

func TestScooterBountyValidUserAdmin(t *testing.T) {
	admin := users.Admin{}

	state := states.ScooterBounty{}
	state.User = admin

	ret, _ := state.AllowedUser()
	if ret != false {
		t.Fatalf("users.Admin expected to NOT be allowed")
	}
}

func TestScooterBountyIsValid(t *testing.T) {
	user := users.User{}

	state := states.ScooterBounty{}
	state.User = user

	ret, _ := state.IsValid()

	if ret != false {
		t.Fatalf("Expected invalid bounty due commom user")
	}
}

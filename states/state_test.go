package states_test

import (
	"github.com/ramonmedeiros/state_machine_go/states"
	"github.com/ramonmedeiros/state_machine_go/users"
	"testing"
)

func TestStateStruct(t *testing.T) {
	state := states.ScooterState{}
	ret, err := state.Next()

	if err == nil {
		t.Fatalf("State expected to not have next function %v", ret)
	}

}

func TestScooterStateIsValid(t *testing.T) {
	state := states.ScooterState{}
	ret, _ := state.IsValid()

	if ret != false {
		t.Fatalf("ScooterState must not have implementation")
	}

}

func TestScooterStateValidUserUser(t *testing.T) {
	user := users.User{}

	state := states.ScooterState{}
	state.User = user

	ret, _ := state.AllowedUser()
	if ret == false {
		t.Fatalf("users.User expected to be allowed")
	}
}

func TestScooterStateValidUserHunter(t *testing.T) {
	hunter := users.Hunter{}

	state := states.ScooterState{}
	state.User = hunter

	ret, _ := state.AllowedUser()
	if ret == false {
		t.Fatalf("users.Hunter expected to be allowed")
	}
}

func TestScooterStateValidUserAdmin(t *testing.T) {
	admin := users.Admin{}

	state := states.ScooterState{}
	state.User = admin

	ret, _ := state.AllowedUser()
	if ret == false {
		t.Fatalf("users.Admin expected to be allowed")
	}
}

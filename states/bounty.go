package states

import (
	"fmt"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"time"
)

// ScooterBounty based on ScooterState
type ScooterBounty struct {
	ScooterState
}

// Next return the next state based on conditions
func (state *ScooterBounty) Next() (interface{}, error) {
	if reflect.TypeOf(state.User) != reflect.TypeOf(users.Hunter{}) {
		return false, fmt.Errorf("Expected Hunter user to change state")
	}

	collected := ScooterCollected{}
	collected.User = state.User
	collected.BatteryLevel = state.BatteryLevel
	collected.LastStateChange = time.Now()

	return collected, nil
}

// IsValid validates the actual state of the struct
func (state *ScooterBounty) IsValid() (bool, error) {

	usersValid, usersMsg := state.AllowedUser()
	if usersValid == false {
		return false, fmt.Errorf("%v", usersMsg)
	}

	return true, nil
}

func (state *ScooterBounty) AllowedUser() (bool, error) {
	allowedUser, _ := state.AllowedUsers()
	for i, _ := range allowedUser {
		if reflect.TypeOf(allowedUser[i]) == reflect.TypeOf(state.User) {
			return true, nil
		}
	}
	return false, fmt.Errorf("User %v not allowed", state.User)
}

func (state *ScooterBounty) AllowedUsers() ([]interface{}, error) {
	return []interface{}{users.Hunter{}, nil}, nil
}

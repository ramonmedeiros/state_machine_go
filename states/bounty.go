package states

import (
	"fmt"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"time"
)

type ScooterBounty struct {
	ScooterState
}

func (state *ScooterBounty) Next() (interface{}, error) {
	if (reflect.TypeOf(state.User) != reflect.TypeOf(&users.Hunter{})) {
		return false, fmt.Errorf("Hunter user is expected")
	}

	collected := ScooterCollected{}
	collected.Name = state.Name
	collected.User = state.User
	collected.BatteryLevel = state.BatteryLevel
	collected.LastStateChange = time.Now()

	return collected, nil
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
	return []interface{}{users.Hunter{}}, nil
}

package states

import (
	"fmt"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"time"
)

type ScooterCollected struct {
	ScooterState
}

func (state *ScooterCollected) Next() (interface{}, error) {
	if (reflect.TypeOf(state.User) != reflect.TypeOf(&users.Hunter{})) {
		return false, fmt.Errorf("Hunter user is expected")
	}

	dropped := ScooterDropped{}
	dropped.Name = state.Name
	dropped.User = state.User
	dropped.BatteryLevel = state.BatteryLevel
	dropped.LastStateChange = time.Now()

	return dropped, nil
}

func (state *ScooterCollected) AllowedUser() (bool, error) {
	allowedUser, _ := state.AllowedUsers()
	for i, _ := range allowedUser {
		if reflect.TypeOf(allowedUser[i]) == reflect.TypeOf(state.User) {
			return true, nil
		}
	}
	return false, fmt.Errorf("User %v not allowed", state.User)
}

func (state *ScooterCollected) AllowedUsers() ([]interface{}, error) {
	return []interface{}{users.Hunter{}}, nil
}

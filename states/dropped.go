package states

import (
	"fmt"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"time"
)

type ScooterDropped struct {
	ScooterState
}

func (state *ScooterDropped) Next() (interface{}, error) {
	if (reflect.TypeOf(state.User) != reflect.TypeOf(&users.Hunter{})) {
		return false, fmt.Errorf("Hunter user is expected")
	}

	dropped := ScooterReady{}
	dropped.User = nil
	dropped.BatteryLevel = 100
	dropped.LastStateChange = time.Now()

	return dropped, nil
}

func (state *ScooterDropped) AllowedUser() (bool, error) {
	allowedUser, _ := state.AllowedUsers()
	for i, _ := range allowedUser {
		if reflect.TypeOf(allowedUser[i]) == reflect.TypeOf(state.User) {
			return true, nil
		}
	}
	return false, fmt.Errorf("User %v not allowed", state.User)
}

func (state *ScooterDropped) AllowedUsers() ([]interface{}, error) {
	return []interface{}{users.Hunter{}}, nil
}

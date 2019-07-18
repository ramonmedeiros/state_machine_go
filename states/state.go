package states

import (
	"fmt"
	"github.com/ramonmedeiros/state_machine_go/users"
	"reflect"
	"time"
)

type ScooterState struct {
	Name            string
	User            interface{}
	BatteryLevel    int
	LastStateChange time.Time
}

func (state *ScooterState) Next() (bool, error) {
	return false, fmt.Errorf("no next state implemented")
}

func (state *ScooterState) IsValid() (bool, error) {
	return false, fmt.Errorf("ScooterState is just an abstraction. Need proper implementation")
}

func (state *ScooterState) AllowedUser(user interface{}) (bool, error) {
	allowedUser, _ := state.AllowedUsers()
	for i, _ := range allowedUser {
		if reflect.TypeOf(allowedUser[i]) == reflect.TypeOf(user) {
			return true, nil
		}
	}
	return false, fmt.Errorf("User %v not allowed", user)
}

func (state *ScooterState) AllowedUsers() ([]interface{}, error) {
	return []interface{}{users.User{}, users.Hunter{}, users.Admin{}}, nil
}

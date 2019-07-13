package states

import (
	"fmt"
	"time"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/users"
)

var AllowedUsers = []interface{}{users.User{}, users.Hunter{}, users.Admin{}}

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
	for i, _ := range AllowedUsers {
        if reflect.TypeOf(AllowedUsers[i]) == reflect.TypeOf(user) {
			return true, nil
		}
	}
	return false, fmt.Errorf("User %v not allowed", user)
}


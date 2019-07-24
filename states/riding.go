package states

import (
	"fmt"
	"time"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/users"
)

type ScooterRiding struct {
	ScooterState
}

func (state *ScooterRiding) Next() (interface{}, error) {
	// low battery: move state
	if state.BatteryLevel < 20 {
		batteryLow := ScooterBatteryLow{}
		batteryLow.User = state.User
		batteryLow.BatteryLevel = state.BatteryLevel
		batteryLow.LastStateChange = time.Now()

		return batteryLow, nil
	}

	// user deattached: back to ride
	if state.User == nil {
		ready := ScooterReady{}
		ready.User = state.User
		ready.BatteryLevel = state.BatteryLevel
		ready.LastStateChange = time.Now()
		return ready, nil
	}

	// no change, keep
	return *state, nil
}

func (state *ScooterRiding) IsValid() (bool, error) {
	if state.BatteryLevel < 20 {
		return false, fmt.Errorf("BatteryLevel too low, should change status")
	}

	usersValid, usersMsg := state.AllowedUser()
	if usersValid == false {
		return false, fmt.Errorf("%v", usersMsg)
	}

	return true, nil
}

func (state *ScooterRiding) AllowedUser() (bool, error) {
	allowedUser, _ := state.AllowedUsers()
	for i, _ := range allowedUser {
		if reflect.TypeOf(allowedUser[i]) == reflect.TypeOf(state.User) {
			return true, nil
		}
	}
	return false, fmt.Errorf("User %v not allowed", state.User)
}

func (state *ScooterRiding) AllowedUsers() ([]interface{}, error) {
	return []interface{}{users.User{}, users.Hunter{}, users.Admin{}}, nil
}

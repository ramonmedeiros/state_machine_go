package states

import (
	"fmt"
	"time"
)

// ScooterBatteryLow based on ScooterState
type ScooterBatteryLow struct {
	ScooterState
}

// Next return the next state based on conditions
func (state *ScooterBatteryLow) Next() (interface{}, error) {
	bounty := ScooterBounty{}
	bounty.User = nil
	bounty.BatteryLevel = state.BatteryLevel
	bounty.LastStateChange = time.Now()
	return bounty, nil

}

// IsValid validates the actual state of the struct
func (state *ScooterBatteryLow) IsValid() (bool, error) {
	// battery should be less than 20%
	if state.BatteryLevel >= 20 {
		return false, fmt.Errorf("BatteryLow requires 20%% of level, %v found", state.BatteryLevel)
	}

	usersValid, usersMsg := state.AllowedUser()
	if usersValid == false {
		return false, fmt.Errorf("%v", usersMsg)
	}

	return true, nil
}

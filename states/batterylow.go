package states

import (
	"fmt"
	"time"
)

type ScooterBatteryLow struct {
	ScooterState
}

func (state *ScooterBatteryLow) Next() (interface{}, error) {
	bounty := ScooterBounty{}
	bounty.Name = state.Name
	bounty.User = nil
	bounty.BatteryLevel = state.BatteryLevel
	bounty.LastStateChange = time.Now()
	return bounty, nil

}

func (state *ScooterBatteryLow) IsValid() (bool, error) {
	// battery should be less than 20%
	if state.BatteryLevel >= 20 {
		return false, fmt.Errorf("BatteryLow requires 20%% of level, %v found", state.BatteryLevel)
	}

	return true, nil
}

package states

import "time"

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
    return true, nil
}


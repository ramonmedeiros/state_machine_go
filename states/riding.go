package states

import (
	"time"
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
	return true, nil
}

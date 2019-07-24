package states

import (
	"fmt"
	"time"
)

var Now = time.Now
var bountyTime = time.Date(Now().Year(), Now().Month(), Now().Day(), 21, 30, 0, 0, Now().Location())

type ScooterReady struct {
	ScooterState
}

func (state *ScooterReady) Next() (interface{}, error) {

	// last changed was more than 48h: go to Unknown
	if state.LastStateChange.Add(time.Hour * 48).Before(Now()) {
		unknown := ScooterUnknown{}
		unknown.User = nil
		unknown.BatteryLevel = state.BatteryLevel
		unknown.LastStateChange = Now()
		return unknown, nil
	}

	// more than 21:30: set as Bounty
	if Now().After(bountyTime) {
		bounty := ScooterBounty{}
		bounty.BatteryLevel = state.BatteryLevel
		bounty.User = state.User
		bounty.LastStateChange = Now()
		return bounty, nil
	}

	// no user attached: stay ready
	if state.User == nil {
		return *state, nil
	}

	// with a user: go to rinding
	riding := ScooterRiding{}
	riding.BatteryLevel = state.BatteryLevel
	riding.User = state.User
	riding.LastStateChange = Now()
	return riding, nil
}

func (state *ScooterReady) IsValid() (bool, error) {
	if state.User != nil {
		return false, fmt.Errorf("Ready state does not expect User")
	}

	if Now().After(bountyTime) {
		return false, fmt.Errorf("After 21:30, state should be ScootterBounty")
	}

	if state.LastStateChange.Add(time.Hour * 48).Before(Now()) {
		return false, fmt.Errorf("48h without change: must be Unknown")
	}

	usersValid, usersMsg := state.AllowedUser()
	if usersValid == false {
		return false, fmt.Errorf("%v", usersMsg)
	}

	return true, nil
}

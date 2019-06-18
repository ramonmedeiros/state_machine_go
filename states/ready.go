package states

import (
    "fmt"
    "time"
)

type ScooterReady struct {
    ScooterState
}

// this will be mocked on tests
func Now() time.Time {
	return time.Now()
}

func (state *ScooterReady) Next() (interface{}, error) {

    // last changed was more than 48h: go to Unknown
    if (state.LastStateChange.Add(time.Hour * 48).Before(Now())) {
        unknown := ScooterUnknown{}
        unknown.Name = state.Name
        unknown.User = nil
        unknown.BatteryLevel = state.BatteryLevel
        unknown.LastStateChange = Now()
        return unknown, nil
    }

    // more than 21:30: set as Bounty
    now := Now()
    bountyTime := time.Date(now.Year(), now.Month(), now.Day(), 21, 30, 0, 0, now.Location())
    if now.After(bountyTime) {
        bounty := ScooterBounty{}
        bounty.Name = state.Name
        bounty.BatteryLevel = state.BatteryLevel
        bounty.User = state.User
        bounty.LastStateChange = Now()
        return bounty, nil
    }

    // next state needs user
    if (state.User == nil) {
        return nil, fmt.Errorf("Next state needs associate user")
    }

    //  go to rinding
    riding := ScooterRiding{}
    riding.Name = state.Name
    riding.BatteryLevel = state.BatteryLevel
    riding.User = state.User
    riding.LastStateChange = Now()
    return riding, nil
}

func (state *ScooterReady) IsValid() (bool, error) {
    return true, nil
}



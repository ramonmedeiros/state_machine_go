package states

import (
    "time"
)

var Now = time.Now

type ScooterReady struct {
    ScooterState
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
    bountyTime := time.Date(Now().Year(), Now().Month(), Now().Day(), 21, 30, 0, 0, Now().Location())
    if Now().After(bountyTime) {
        bounty := ScooterBounty{}
        bounty.Name = state.Name
        bounty.BatteryLevel = state.BatteryLevel
        bounty.User = state.User
        bounty.LastStateChange = Now()
        return bounty, nil
    }

    // no user attached: stay ready
    if (state.User == nil) {
        return *state, nil
    }

    // with a user: go to rinding
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



package states_test

import (
    "time"
    "testing"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/states"
)

var bountyTime = func() time.Time { return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 30, 1, 0, time.Now().Location()) }

func TestReadyToUnknownState(t *testing.T) {
    twodaysago := time.Now().Add(time.Hour * -48)
    state := states.ScooterReady{}
    state.Name = "test-name"
    state.User = nil
    state.BatteryLevel = 100
    state.LastStateChange = twodaysago

    newstate, _ := state.Next()

    if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterUnknown{})) {
        t.Fatalf("Expected unknow, found %v", reflect.TypeOf(newstate))
    }
}

func TestBountyState(t *testing.T) {
    state := states.ScooterReady{}
    state.Name = "test-name"
    state.User = nil
    state.BatteryLevel = 100
    state.LastStateChange = time.Now()

    // mock time
    states.Now = bountyTime

    newstate, _ := state.Next()

    if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterBounty{})) {
        t.Fatalf("Expected unknow, found %v", reflect.TypeOf(newstate))
    }
}


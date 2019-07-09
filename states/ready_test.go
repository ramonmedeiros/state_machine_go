package states_test

import (
    "time"
    "testing"
    "reflect"
    "github.com/ramonmedeiros/state_machine_go/states"
)


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

    // mock time to 21:30
    states.Now = func() time.Time { return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 30, 1, 0, time.Now().Location()) }
    newstate, _ := state.Next()

    // rollback mock
    states.Now = time.Now

    if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterBounty{})) {
        t.Fatalf("Expected Bounty, found %v", reflect.TypeOf(newstate))
    }
}

func TestKeepReady(t *testing.T) {
    state := states.ScooterReady{}
    state.Name = "test-name"
    state.User = nil
    state.BatteryLevel = 100
    state.LastStateChange = time.Now()

    newstate, _ := state.Next()

    if (reflect.TypeOf(newstate) != reflect.TypeOf(states.ScooterReady{})) {
        t.Fatalf("Expected Ready, found %v", reflect.TypeOf(newstate))
    }
}
